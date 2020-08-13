// Code generated by go-bindata.
// sources:
// data/github_auth.tmpl
// data/github_auth_result.tmpl
// data/google_auth.tmpl
// DO NOT EDIT!

package authn

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _dataGithub_authTmpl = []byte(`<!doctype html>

<html>
<head>
  <meta charset="utf-8">
  <title>Docker Registry Authentication</title>
  <style>
    .github-icon {
      background-image: url('data:image/svg+xml;utf8,<svg version="1.1" id="Github" xmlns="http://www.w3.org/2000/svg" x="0px" y="0px"  width="155px" height="155px" viewBox="0 0 155 155" enable-background="new 0 0 155 155" xml:space="preserve"><path id="path26" fill="#cccccc" d="M78.012,3.04c-41.685,0-75.486,33.794-75.486,75.486c0,33.35,21.629,61.645,51.622,71.625 c3.772,0.699,5.157-1.637,5.157-3.631c0-1.799-0.07-7.746-0.103-14.054c-21,4.566-25.431-8.906-25.431-8.906  c-3.434-8.724-8.381-11.045-8.381-11.045c-6.849-4.685,0.517-4.589,0.517-4.589c7.58,0.533,11.571,7.78,11.571,7.78 c6.732,11.538,17.659,8.202,21.966,6.273c0.678-4.878,2.634-8.208,4.793-10.094c-16.767-1.907-34.392-8.382-34.392-37.305 c0-8.24,2.949-14.975,7.778-20.261c-0.784-1.901-3.368-9.579,0.731-19.977c0,0,6.34-2.027,20.764,7.739 c6.021-1.672,12.479-2.511,18.895-2.541c6.414,0.03,12.877,0.869,18.909,2.541c14.407-9.766,20.737-7.739,20.737-7.739  c4.109,10.397,1.525,18.075,0.742,19.977c4.84,5.286,7.768,12.021,7.768,20.261c0,28.993-17.659,35.376-34.469,37.245 c2.708,2.342,5.121,6.936,5.121,13.979c0,10.1-0.087,18.229-0.087,20.715c0,2.01,1.358,4.363,5.185,3.623 c29.976-9.993,51.578-38.277,51.578-71.617C153.496,36.833,119.699,3.04,78.012,3.04"/></svg>');
      background-repeat: no-repeat;
      background-size: contain;
      display: inline-block;
      width: 25px;
      height: 25px;
    }
    body {
      color: #000;
      background: #fff;
      font-family: sans-serif;
      padding: 4em 4em;
    }
    a { color: #000; }
    #panel p { text-align:center; }
    #login-with-github {
      color: #fff;
      background-color: #2a2a2a;
      font-size: 1em;
      text-decoration: none;
      line-height: 3em;
      padding: 0.5em 2em 0.5em 2em;
      display: inline-block;
      height: 3em;
      border-radius: 0.75em;
      cursor: pointer;
      transition: all 0.25s;
    }
    #login-with-github:hover {
      background-color: #444444;
    }
    #login-with-github:active {
      background-color: #101010;
    }
    #login-with-github .github-icon {
      margin: 0 0.5em 0 0;
      position: relative;
      top: 5px;
    }
    #login-with-github code {
      font-size: 1.4em;
    }
    #revoke-access {
      color: #666;
      font-size: 0.8em;
      text-decoration: none;
    }
    #revoke-access:hover {
      text-decoration: underline;
    }
  </style>
</head>

<body>
  <div id="panel">
    <p>
      <a id="login-with-github" href="{{.GithubWebUri}}/login/oauth/authorize?scope=user:email%20read:org&client_id={{.ClientId}}">
        <i class="github-icon"></i>
        Login{{if .Organization}} to <code>@{{.Organization}}</code>{{end}} with GitHub
      </a>
    </p>
    <p>
      <a id="revoke-access" href="{{.GithubWebUri}}/settings/applications">Revoke access</a>
    </p>
  </div>
</body>
</html>
`)

func dataGithub_authTmplBytes() ([]byte, error) {
	return _dataGithub_authTmpl, nil
}

func dataGithub_authTmpl() (*asset, error) {
	bytes, err := dataGithub_authTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/github_auth.tmpl", size: 2946, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataGithub_auth_resultTmpl = []byte(`<!doctype html>

<html>
<head>
  <meta charset="utf-8">
  <title>Docker Registry Authentication</title>
  <style>
    body {
      color: #000;
      background: #fff;
      font-family: sans-serif;
      padding: 4em 4em;
    }
    hr {
      border: none;
      border-top: 1px solid #ccc;
    }
    .message code {
      font-size: 1.4em;
      background: #ccc;
      border-radius: 0.5em;
      padding: 0.25em 0.5em;
      margin: 0 0.25em 0 0.25em;
    }
    .command {
      font-size: 2em;
      line-height: 2em;
      color: #222;
      background: #fafafa;
      padding: 1em 1em 1.2em 1em;
      margin: 1em 0;
      border-radius: 0.5em;
      text-shadow: 0px 1px 0px #fff;
    }
    .command span {
      user-select: none;
      -moz-user-select: none;
      -webkit-user-select: none;
      -ms-user-select: none;
    }
  </style>
</head>
<body>
  <p class="message">
    You are successfully authenticated for the Docker Registry{{if .Organization}} with the <code>@{{.Organization}}</code> Github organization{{end}}.
    Use the following username and password to login into the registry:
  </p>
  <hr>
  <pre class="command"><span>$ </span>cellery login -u {{.Username}} -p {{.Password}} {{if .RegistryUrl}}{{.RegistryUrl}}{{else}}docker.example.com{{end}}</pre>
</body>
</html>
`)

func dataGithub_auth_resultTmplBytes() ([]byte, error) {
	return _dataGithub_auth_resultTmpl, nil
}

func dataGithub_auth_resultTmpl() (*asset, error) {
	bytes, err := dataGithub_auth_resultTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/github_auth_result.tmpl", size: 1300, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataGoogle_authTmpl = []byte(`<html itemscope itemtype="http://schema.org/Article">
<head>
  <script src="//ajax.googleapis.com/ajax/libs/jquery/1.8.2/jquery.min.js"></script>
  <script src="https://apis.google.com/js/client:platform.js?onload=start" async defer></script>
  <script>
		function checkLogin() {
			var auth2 = gapi.auth2.getAuthInstance();
			if (auth2.isSignedIn.get()) {
				$('#result').text('validating existing token...');
				var id_token = auth2.currentUser.get().getAuthResponse().id_token;
				$.ajax({
					type: 'POST',
					url: '/google_auth',
					contentType: 'application/json; charset=utf-8',
					processData: false,
					data: JSON.stringify({'action': 'check', 'token': id_token}),
					success: function(result) {
						$('#result').text(result);
					},
					error: function(xhr) {
						$('#result').text('error: ' + xhr.responseText);
					},
				});
			} else {
				console.log('not logged in');
			}
		}
    function start() {
      gapi.load('auth2', function() {
				gapi.auth2.init({client_id: '{{.ClientId}}'}).then(checkLogin);
			});
    }
  </script>
</head>
<body>
	<button id="signinButton">Sign in with Google</button>
	<script>
		$('#signinButton').click(function() {
			// signInCallback defined in step 6.
			var auth2 = gapi.auth2.getAuthInstance();
			auth2.grantOfflineAccess({'redirect_uri': 'postmessage', 'prompt': 'consent'}).then(function(authResult) {
				console.log(authResult);
				$.ajax({
					type: 'POST',
					url: '/google_auth',
					contentType: 'application/json; charset=utf-8',
					processData: false,
					data: JSON.stringify({'action': 'sign_in', 'code': authResult['code']}),
					success: function(result) {
						$('#result').text(result);
						console.log("result:", result);
					},
					error: function(xhr) {
						$('#result').text('error: ' + xhr.responseText);
						console.log('error:', xhr.responseText);
					},
				});
			});
		});
	</script>
	<button id="signOutButton">Sign out</button>
	<script>
		$('#signOutButton').click(function() {
			var auth2 = gapi.auth2.getAuthInstance();
			if (auth2.isSignedIn.get()) {
				$('#result').text('validating existing token...');
				var id_token = auth2.currentUser.get().getAuthResponse().id_token;
				// Perform server-side sign out.
				$.ajax({
					type: 'POST',
					url: '/google_auth',
					contentType: 'application/json; charset=utf-8',
					data: JSON.stringify({'action': 'sign_out', 'token': id_token}),
					processData: false,
					success: function() {},
					error: function() {},
					complete: function(xhr) {
						console.log('sign out result:', xhr.responseText);
						gapi.auth2.getAuthInstance().disconnect();
						$('#result').text('signed out');
					},
				});
			} else {
				$('#result').text('not logged in');
			}
		});
	</script>
	<div id="result"></div>
</body>
</html>
`)

func dataGoogle_authTmplBytes() ([]byte, error) {
	return _dataGoogle_authTmpl, nil
}

func dataGoogle_authTmpl() (*asset, error) {
	bytes, err := dataGoogle_authTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/google_auth.tmpl", size: 2817, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"data/github_auth.tmpl": dataGithub_authTmpl,
	"data/github_auth_result.tmpl": dataGithub_auth_resultTmpl,
	"data/google_auth.tmpl": dataGoogle_authTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"data": &bintree{nil, map[string]*bintree{
		"github_auth.tmpl": &bintree{dataGithub_authTmpl, map[string]*bintree{}},
		"github_auth_result.tmpl": &bintree{dataGithub_auth_resultTmpl, map[string]*bintree{}},
		"google_auth.tmpl": &bintree{dataGoogle_authTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

