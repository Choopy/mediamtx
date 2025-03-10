package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/bluenviron/mediamtx/internal/conf"
	"github.com/bluenviron/mediamtx/internal/logger"
	"github.com/stretchr/testify/require"
)

type testParent struct{}

func (testParent) Log(_ logger.Level, _ string, _ ...interface{}) {
}

func (testParent) APIConfigSet(_ *conf.Conf) {}

func writeTempFile(byts []byte) (string, error) {
	tmpf, err := os.CreateTemp(os.TempDir(), "rtsp-")
	if err != nil {
		return "", err
	}
	defer tmpf.Close()

	_, err = tmpf.Write(byts)
	if err != nil {
		return "", err
	}

	return tmpf.Name(), nil
}

func tempConf(t *testing.T, cnt string) *conf.Conf {
	fi, err := writeTempFile([]byte(cnt))
	require.NoError(t, err)
	defer os.Remove(fi)

	cnf, _, err := conf.Load(fi, nil)
	require.NoError(t, err)

	return cnf
}

func httpRequest(t *testing.T, hc *http.Client, method string, ur string, in interface{}, out interface{}) {
	buf := func() io.Reader {
		if in == nil {
			return nil
		}

		byts, err := json.Marshal(in)
		require.NoError(t, err)

		return bytes.NewBuffer(byts)
	}()

	req, err := http.NewRequest(method, ur, buf)
	require.NoError(t, err)

	res, err := hc.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("bad status code: %d", res.StatusCode)
	}

	if out == nil {
		return
	}

	err = json.NewDecoder(res.Body).Decode(out)
	require.NoError(t, err)
}

func checkError(t *testing.T, msg string, body io.Reader) {
	var resErr map[string]interface{}
	err := json.NewDecoder(body).Decode(&resErr)
	require.NoError(t, err)
	require.Equal(t, map[string]interface{}{"error": msg}, resErr)
}

func TestPaginate(t *testing.T) {
	items := make([]int, 5)
	for i := 0; i < 5; i++ {
		items[i] = i
	}

	pageCount, err := paginate(&items, "1", "1")
	require.NoError(t, err)
	require.Equal(t, 5, pageCount)
	require.Equal(t, []int{1}, items)

	items = make([]int, 5)
	for i := 0; i < 5; i++ {
		items[i] = i
	}

	pageCount, err = paginate(&items, "3", "2")
	require.NoError(t, err)
	require.Equal(t, 2, pageCount)
	require.Equal(t, []int{}, items)

	items = make([]int, 6)
	for i := 0; i < 6; i++ {
		items[i] = i
	}

	pageCount, err = paginate(&items, "4", "1")
	require.NoError(t, err)
	require.Equal(t, 2, pageCount)
	require.Equal(t, []int{4, 5}, items)
}

func TestConfigGlobalGet(t *testing.T) {
	cnf := tempConf(t, "api: yes\n")

	api := API{
		Address:     "localhost:9997",
		ReadTimeout: conf.StringDuration(10 * time.Second),
		Conf:        cnf,
		Parent:      &testParent{},
	}
	err := api.Initialize()
	require.NoError(t, err)
	defer api.Close()

	hc := &http.Client{Transport: &http.Transport{}}

	var out map[string]interface{}
	httpRequest(t, hc, http.MethodGet, "http://localhost:9997/v3/config/global/get", nil, &out)
	require.Equal(t, true, out["api"])
}

func TestConfigGlobalPatch(t *testing.T) {
	cnf := tempConf(t, "api: yes\n")

	api := API{
		Address:     "localhost:9997",
		ReadTimeout: conf.StringDuration(10 * time.Second),
		Conf:        cnf,
		Parent:      &testParent{},
	}
	err := api.Initialize()
	require.NoError(t, err)
	defer api.Close()

	hc := &http.Client{Transport: &http.Transport{}}

	httpRequest(t, hc, http.MethodPatch, "http://localhost:9997/v3/config/global/patch", map[string]interface{}{
		"rtmp":            false,
		"readTimeout":     "7s",
		"protocols":       []string{"tcp"},
		"readBufferCount": 4096, // test setting a deprecated parameter
	}, nil)

	time.Sleep(500 * time.Millisecond)

	var out map[string]interface{}
	httpRequest(t, hc, http.MethodGet, "http://localhost:9997/v3/config/global/get", nil, &out)
	require.Equal(t, false, out["rtmp"])
	require.Equal(t, "7s", out["readTimeout"])
	require.Equal(t, []interface{}{"tcp"}, out["protocols"])
	require.Equal(t, float64(4096), out["readBufferCount"])
}

func TestAPIConfigGlobalPatchUnknownField(t *testing.T) { //nolint:dupl
	cnf := tempConf(t, "api: yes\n")

	api := API{
		Address:     "localhost:9997",
		ReadTimeout: conf.StringDuration(10 * time.Second),
		Conf:        cnf,
		Parent:      &testParent{},
	}
	err := api.Initialize()
	require.NoError(t, err)
	defer api.Close()

	b := map[string]interface{}{
		"test": "asd",
	}

	byts, err := json.Marshal(b)
	require.NoError(t, err)

	hc := &http.Client{Transport: &http.Transport{}}

	req, err := http.NewRequest(http.MethodPatch, "http://localhost:9997/v3/config/global/patch", bytes.NewReader(byts))
	require.NoError(t, err)

	res, err := hc.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	require.Equal(t, http.StatusBadRequest, res.StatusCode)
	checkError(t, "json: unknown field \"test\"", res.Body)
}

func TestAPIConfigPathDefaultsGet(t *testing.T) {
	cnf := tempConf(t, "api: yes\n")

	api := API{
		Address:     "localhost:9997",
		ReadTimeout: conf.StringDuration(10 * time.Second),
		Conf:        cnf,
		Parent:      &testParent{},
	}
	err := api.Initialize()
	require.NoError(t, err)
	defer api.Close()

	hc := &http.Client{Transport: &http.Transport{}}

	var out map[string]interface{}
	httpRequest(t, hc, http.MethodGet, "http://localhost:9997/v3/config/pathdefaults/get", nil, &out)
	require.Equal(t, "publisher", out["source"])
}

func TestAPIConfigPathDefaultsPatch(t *testing.T) {
	cnf := tempConf(t, "api: yes\n")

	api := API{
		Address:     "localhost:9997",
		ReadTimeout: conf.StringDuration(10 * time.Second),
		Conf:        cnf,
		Parent:      &testParent{},
	}
	err := api.Initialize()
	require.NoError(t, err)
	defer api.Close()

	hc := &http.Client{Transport: &http.Transport{}}

	httpRequest(t, hc, http.MethodPatch, "http://localhost:9997/v3/config/pathdefaults/patch", map[string]interface{}{
		"readUser": "myuser",
		"readPass": "mypass",
	}, nil)

	time.Sleep(500 * time.Millisecond)

	var out map[string]interface{}
	httpRequest(t, hc, http.MethodGet, "http://localhost:9997/v3/config/pathdefaults/get", nil, &out)
	require.Equal(t, "myuser", out["readUser"])
	require.Equal(t, "mypass", out["readPass"])
}

func TestAPIConfigPathsList(t *testing.T) {
	cnf := tempConf(t, "api: yes\n"+
		"paths:\n"+
		"  path1:\n"+
		"    readUser: myuser1\n"+
		"    readPass: mypass1\n"+
		"  path2:\n"+
		"    readUser: myuser2\n"+
		"    readPass: mypass2\n")

	api := API{
		Address:     "localhost:9997",
		ReadTimeout: conf.StringDuration(10 * time.Second),
		Conf:        cnf,
		Parent:      &testParent{},
	}
	err := api.Initialize()
	require.NoError(t, err)
	defer api.Close()

	type pathConfig map[string]interface{}

	type listRes struct {
		ItemCount int          `json:"itemCount"`
		PageCount int          `json:"pageCount"`
		Items     []pathConfig `json:"items"`
	}

	hc := &http.Client{Transport: &http.Transport{}}

	var out listRes
	httpRequest(t, hc, http.MethodGet, "http://localhost:9997/v3/config/paths/list", nil, &out)
	require.Equal(t, 2, out.ItemCount)
	require.Equal(t, 1, out.PageCount)
	require.Equal(t, "path1", out.Items[0]["name"])
	require.Equal(t, "myuser1", out.Items[0]["readUser"])
	require.Equal(t, "mypass1", out.Items[0]["readPass"])
	require.Equal(t, "path2", out.Items[1]["name"])
	require.Equal(t, "myuser2", out.Items[1]["readUser"])
	require.Equal(t, "mypass2", out.Items[1]["readPass"])
}

func TestAPIConfigPathsGet(t *testing.T) {
	cnf := tempConf(t, "api: yes\n"+
		"paths:\n"+
		"  my/path:\n"+
		"    readUser: myuser\n"+
		"    readPass: mypass\n")

	api := API{
		Address:     "localhost:9997",
		ReadTimeout: conf.StringDuration(10 * time.Second),
		Conf:        cnf,
		Parent:      &testParent{},
	}
	err := api.Initialize()
	require.NoError(t, err)
	defer api.Close()

	hc := &http.Client{Transport: &http.Transport{}}

	var out map[string]interface{}
	httpRequest(t, hc, http.MethodGet, "http://localhost:9997/v3/config/paths/get/my/path", nil, &out)
	require.Equal(t, "my/path", out["name"])
	require.Equal(t, "myuser", out["readUser"])
}

func TestAPIConfigPathsAdd(t *testing.T) {
	cnf := tempConf(t, "api: yes\n")

	api := API{
		Address:     "localhost:9997",
		ReadTimeout: conf.StringDuration(10 * time.Second),
		Conf:        cnf,
		Parent:      &testParent{},
	}
	err := api.Initialize()
	require.NoError(t, err)
	defer api.Close()

	hc := &http.Client{Transport: &http.Transport{}}

	httpRequest(t, hc, http.MethodPost, "http://localhost:9997/v3/config/paths/add/my/path", map[string]interface{}{
		"source":                   "rtsp://127.0.0.1:9999/mypath",
		"sourceOnDemand":           true,
		"disablePublisherOverride": true, // test setting a deprecated parameter
		"rpiCameraVFlip":           true,
	}, nil)

	var out map[string]interface{}
	httpRequest(t, hc, http.MethodGet, "http://localhost:9997/v3/config/paths/get/my/path", nil, &out)
	require.Equal(t, "rtsp://127.0.0.1:9999/mypath", out["source"])
	require.Equal(t, true, out["sourceOnDemand"])
	require.Equal(t, true, out["disablePublisherOverride"])
	require.Equal(t, true, out["rpiCameraVFlip"])
}

func TestAPIConfigPathsAddUnknownField(t *testing.T) { //nolint:dupl
	cnf := tempConf(t, "api: yes\n")

	api := API{
		Address:     "localhost:9997",
		ReadTimeout: conf.StringDuration(10 * time.Second),
		Conf:        cnf,
		Parent:      &testParent{},
	}
	err := api.Initialize()
	require.NoError(t, err)
	defer api.Close()

	b := map[string]interface{}{
		"test": "asd",
	}

	byts, err := json.Marshal(b)
	require.NoError(t, err)

	hc := &http.Client{Transport: &http.Transport{}}

	req, err := http.NewRequest(http.MethodPost,
		"http://localhost:9997/v3/config/paths/add/my/path", bytes.NewReader(byts))
	require.NoError(t, err)

	res, err := hc.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	require.Equal(t, http.StatusBadRequest, res.StatusCode)
	checkError(t, "json: unknown field \"test\"", res.Body)
}

func TestAPIConfigPathsPatch(t *testing.T) { //nolint:dupl
	cnf := tempConf(t, "api: yes\n")

	api := API{
		Address:     "localhost:9997",
		ReadTimeout: conf.StringDuration(10 * time.Second),
		Conf:        cnf,
		Parent:      &testParent{},
	}
	err := api.Initialize()
	require.NoError(t, err)
	defer api.Close()

	hc := &http.Client{Transport: &http.Transport{}}

	httpRequest(t, hc, http.MethodPost, "http://localhost:9997/v3/config/paths/add/my/path", map[string]interface{}{
		"source":                   "rtsp://127.0.0.1:9999/mypath",
		"sourceOnDemand":           true,
		"disablePublisherOverride": true, // test setting a deprecated parameter
		"rpiCameraVFlip":           true,
	}, nil)

	httpRequest(t, hc, http.MethodPatch, "http://localhost:9997/v3/config/paths/patch/my/path", map[string]interface{}{
		"source":         "rtsp://127.0.0.1:9998/mypath",
		"sourceOnDemand": true,
	}, nil)

	var out map[string]interface{}
	httpRequest(t, hc, http.MethodGet, "http://localhost:9997/v3/config/paths/get/my/path", nil, &out)
	require.Equal(t, "rtsp://127.0.0.1:9998/mypath", out["source"])
	require.Equal(t, true, out["sourceOnDemand"])
	require.Equal(t, true, out["disablePublisherOverride"])
	require.Equal(t, true, out["rpiCameraVFlip"])
}

func TestAPIConfigPathsReplace(t *testing.T) { //nolint:dupl
	cnf := tempConf(t, "api: yes\n")

	api := API{
		Address:     "localhost:9997",
		ReadTimeout: conf.StringDuration(10 * time.Second),
		Conf:        cnf,
		Parent:      &testParent{},
	}
	err := api.Initialize()
	require.NoError(t, err)
	defer api.Close()

	hc := &http.Client{Transport: &http.Transport{}}

	httpRequest(t, hc, http.MethodPost, "http://localhost:9997/v3/config/paths/add/my/path", map[string]interface{}{
		"source":                   "rtsp://127.0.0.1:9999/mypath",
		"sourceOnDemand":           true,
		"disablePublisherOverride": true, // test setting a deprecated parameter
		"rpiCameraVFlip":           true,
	}, nil)

	httpRequest(t, hc, http.MethodPost, "http://localhost:9997/v3/config/paths/replace/my/path", map[string]interface{}{
		"source":         "rtsp://127.0.0.1:9998/mypath",
		"sourceOnDemand": true,
	}, nil)

	var out map[string]interface{}
	httpRequest(t, hc, http.MethodGet, "http://localhost:9997/v3/config/paths/get/my/path", nil, &out)
	require.Equal(t, "rtsp://127.0.0.1:9998/mypath", out["source"])
	require.Equal(t, true, out["sourceOnDemand"])
	require.Equal(t, nil, out["disablePublisherOverride"])
	require.Equal(t, false, out["rpiCameraVFlip"])
}

func TestAPIConfigPathsDelete(t *testing.T) {
	cnf := tempConf(t, "api: yes\n")

	api := API{
		Address:     "localhost:9997",
		ReadTimeout: conf.StringDuration(10 * time.Second),
		Conf:        cnf,
		Parent:      &testParent{},
	}
	err := api.Initialize()
	require.NoError(t, err)
	defer api.Close()

	hc := &http.Client{Transport: &http.Transport{}}

	httpRequest(t, hc, http.MethodPost, "http://localhost:9997/v3/config/paths/add/my/path", map[string]interface{}{
		"source":         "rtsp://127.0.0.1:9999/mypath",
		"sourceOnDemand": true,
	}, nil)

	httpRequest(t, hc, http.MethodDelete, "http://localhost:9997/v3/config/paths/delete/my/path", nil, nil)

	req, err := http.NewRequest(http.MethodGet, "http://localhost:9997/v3/config/paths/get/my/path", nil)
	require.NoError(t, err)

	res, err := hc.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	require.Equal(t, http.StatusNotFound, res.StatusCode)
	checkError(t, "path configuration not found", res.Body)
}

func TestRecordingsList(t *testing.T) {
	dir, err := os.MkdirTemp("", "mediamtx-playback")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	cnf := tempConf(t, "pathDefaults:\n"+
		"  recordPath: "+filepath.Join(dir, "%path/%Y-%m-%d_%H-%M-%S-%f")+"\n"+
		"paths:\n"+
		"  all_others:\n")

	api := API{
		Address:     "localhost:9997",
		ReadTimeout: conf.StringDuration(10 * time.Second),
		Conf:        cnf,
		Parent:      &testParent{},
	}
	err = api.Initialize()
	require.NoError(t, err)
	defer api.Close()

	err = os.Mkdir(filepath.Join(dir, "mypath1"), 0o755)
	require.NoError(t, err)

	err = os.Mkdir(filepath.Join(dir, "mypath2"), 0o755)
	require.NoError(t, err)

	err = os.WriteFile(filepath.Join(dir, "mypath1", "2008-11-07_11-22-00-000000.mp4"), []byte(""), 0o644)
	require.NoError(t, err)

	err = os.WriteFile(filepath.Join(dir, "mypath1", "2009-11-07_11-22-00-000000.mp4"), []byte(""), 0o644)
	require.NoError(t, err)

	err = os.WriteFile(filepath.Join(dir, "mypath2", "2009-11-07_11-22-00-000000.mp4"), []byte(""), 0o644)
	require.NoError(t, err)

	hc := &http.Client{Transport: &http.Transport{}}

	var out interface{}
	httpRequest(t, hc, http.MethodGet, "http://localhost:9997/v3/recordings/list", nil, &out)
	require.Equal(t, map[string]interface{}{
		"itemCount": float64(2),
		"pageCount": float64(1),
		"items": []interface{}{
			map[string]interface{}{
				"name": "mypath1",
				"segments": []interface{}{
					map[string]interface{}{
						"start": time.Date(2008, 11, 0o7, 11, 22, 0, 0, time.Local).Format(time.RFC3339),
					},
					map[string]interface{}{
						"start": time.Date(2009, 11, 0o7, 11, 22, 0, 0, time.Local).Format(time.RFC3339),
					},
				},
			},
			map[string]interface{}{
				"name": "mypath2",
				"segments": []interface{}{
					map[string]interface{}{
						"start": time.Date(2009, 11, 0o7, 11, 22, 0, 0, time.Local).Format(time.RFC3339),
					},
				},
			},
		},
	}, out)
}

func TestRecordingsGet(t *testing.T) {
	dir, err := os.MkdirTemp("", "mediamtx-playback")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	cnf := tempConf(t, "pathDefaults:\n"+
		"  recordPath: "+filepath.Join(dir, "%path/%Y-%m-%d_%H-%M-%S-%f")+"\n"+
		"paths:\n"+
		"  all_others:\n")

	api := API{
		Address:     "localhost:9997",
		ReadTimeout: conf.StringDuration(10 * time.Second),
		Conf:        cnf,
		Parent:      &testParent{},
	}
	err = api.Initialize()
	require.NoError(t, err)
	defer api.Close()

	err = os.Mkdir(filepath.Join(dir, "mypath1"), 0o755)
	require.NoError(t, err)

	err = os.WriteFile(filepath.Join(dir, "mypath1", "2008-11-07_11-22-00-000000.mp4"), []byte(""), 0o644)
	require.NoError(t, err)

	err = os.WriteFile(filepath.Join(dir, "mypath1", "2009-11-07_11-22-00-000000.mp4"), []byte(""), 0o644)
	require.NoError(t, err)

	hc := &http.Client{Transport: &http.Transport{}}

	var out interface{}
	httpRequest(t, hc, http.MethodGet, "http://localhost:9997/v3/recordings/get/mypath1", nil, &out)
	require.Equal(t, map[string]interface{}{
		"name": "mypath1",
		"segments": []interface{}{
			map[string]interface{}{
				"start": time.Date(2008, 11, 0o7, 11, 22, 0, 0, time.Local).Format(time.RFC3339),
			},
			map[string]interface{}{
				"start": time.Date(2009, 11, 0o7, 11, 22, 0, 0, time.Local).Format(time.RFC3339),
			},
		},
	}, out)
}

func TestRecordingsDeleteSegment(t *testing.T) {
	dir, err := os.MkdirTemp("", "mediamtx-playback")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	cnf := tempConf(t, "pathDefaults:\n"+
		"  recordPath: "+filepath.Join(dir, "%path/%Y-%m-%d_%H-%M-%S-%f")+"\n"+
		"paths:\n"+
		"  all_others:\n")

	api := API{
		Address:     "localhost:9997",
		ReadTimeout: conf.StringDuration(10 * time.Second),
		Conf:        cnf,
		Parent:      &testParent{},
	}
	err = api.Initialize()
	require.NoError(t, err)
	defer api.Close()

	err = os.Mkdir(filepath.Join(dir, "mypath1"), 0o755)
	require.NoError(t, err)

	err = os.WriteFile(filepath.Join(dir, "mypath1", "2008-11-07_11-22-00-000000.mp4"), []byte(""), 0o644)
	require.NoError(t, err)

	hc := &http.Client{Transport: &http.Transport{}}

	v := url.Values{}
	v.Set("path", "mypath1")
	v.Set("start", time.Date(2008, 11, 0o7, 11, 22, 0, 0, time.Local).Format(time.RFC3339))

	u := &url.URL{
		Scheme:   "http",
		Host:     "localhost:9997",
		Path:     "/v3/recordings/deletesegment",
		RawQuery: v.Encode(),
	}

	req, err := http.NewRequest(http.MethodDelete, u.String(), nil)
	require.NoError(t, err)

	res, err := hc.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()
	require.Equal(t, http.StatusOK, res.StatusCode)
}
