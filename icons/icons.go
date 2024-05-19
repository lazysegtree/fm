package icons

/*
	Pulled from https://github.com/acarl005/ls-go
*/

import (
	"path/filepath"
	"strings"
)

type IconStyle struct {
	Icon  string
	Color string
}

func GetElementIcon(file string, isDir bool) IconStyle {
	ext := strings.TrimPrefix(filepath.Ext(file), ".")
	name := file

	if isDir {
		icon := folders["folder"]
		betterIcon, hasBetterIcon := folders[name]
		if hasBetterIcon {
			icon = betterIcon
		}
		return icon
	} else {
		// default icon for all files. try to find a better one though...
		icon := icons["file"]
		// resolve aliased extensions
		extKey := strings.ToLower(ext)
		alias, hasAlias := aliases[extKey]
		if hasAlias {
			extKey = alias
		}

		// see if we can find a better icon based on extension alone
		betterIcon, hasBetterIcon := icons[extKey]
		if hasBetterIcon {
			icon = betterIcon
		}

		// now look for icons based on full names
		fullName := name

		fullName = strings.ToLower(fullName)
		fullAlias, hasFullAlias := aliases[fullName]
		if hasFullAlias {
			fullName = fullAlias
		}
		bestIcon, hasBestIcon := icons[fullName]
		if hasBestIcon {
			icon = bestIcon
		}
		if icon.Color == "NONE" {
			return IconStyle{
				Icon:  icon.Icon,
				Color: "#E5C287",
			}
		}
		return icon
	}
}

var icons = map[string]IconStyle{
	"ai": {
		Icon:  "",
		Color: "#ce6f14",
	},
	"android":      {Icon: "", Color: "#a7c83f"},
	"apple":        {Icon: "", Color: "#78909c"},
	"asm":          {Icon: "󰘚", Color: "#ff7844"},
	"audio":        {Icon: "", Color: "#ee524f"},
	"binary":       {Icon: "", Color: "#ff7844"},
	"c":            {Icon: "", Color: "#0188d2"},
	"cfg":          {Icon: "", Color: "#8B8B8B"},
	"clj":          {Icon: "", Color: "#68b338"},
	"conf":         {Icon: "", Color: "#8B8B8B"},
	"cpp":          {Icon: "", Color: "#0188d2"},
	"css":          {Icon: "", Color: "#2d53e5"},
	"dart":         {Icon: "", Color: "#03589b"},
	"db":           {Icon: "", Color: "#FF8400"},
	"deb":          {Icon: "", Color: "#ab0836"},
	"doc":          {Icon: "", Color: "#295394"},
	"dockerfile":   {Icon: "󰡨", Color: "#099cec"},
	"ebook":        {Icon: "", Color: "#67b500"},
	"env":          {Icon: "", Color: "#eed645"},
	"f":            {Icon: "󱈚", Color: "#8e44ad"},
	"file":         {Icon: "\uf15b", Color: "NONE"},
	"font":         {Icon: "\uf031", Color: "#3498db"},
	"fs":           {Icon: "\ue7a7", Color: "#2ecc71"},
	"gb":           {Icon: "\ue272", Color: "#f1c40f"},
	"gform":        {Icon: "\uf298", Color: "#9b59b6"},
	"git":          {Icon: "\ue702", Color: "#e67e22"},
	"go":           {Icon: "", Color: "#6ed8e5"},
	"graphql":      {Icon: "\ue662", Color: "#e74c3c"},
	"glp":          {Icon: "󰆧", Color: "#3498db"},
	"groovy":       {Icon: "\ue775", Color: "#2ecc71"},
	"gruntfile.js": {Icon: "\ue74c", Color: "#3498db"},
	"gulpfile.js":  {Icon: "\ue610", Color: "#e67e22"},
	"gv":           {Icon: "\ue225", Color: "#9b59b6"},
	"h":            {Icon: "\uf0fd", Color: "#3498db"},
	"haml":         {Icon: "\ue664", Color: "#9b59b6"},
	"hs":           {Icon: "\ue777", Color: "#2980b9"},
	"html":         {Icon: "\uf13b", Color: "#e67e22"},
	"hx":           {Icon: "\ue666", Color: "#e74c3c"},
	"ics":          {Icon: "\uf073", Color: "#f1c40f"},
	"image":        {Icon: "\uf1c5", Color: "#e74c3c"},
	"iml":          {Icon: "\ue7b5", Color: "#3498db"},
	"ini":          {Icon: "󰅪", Color: "#f1c40f"},
	"ino":          {Icon: "\ue255", Color: "#2ecc71"},
	"iso":          {Icon: "󰋊", Color: "#f1c40f"},
	"jade":         {Icon: "\ue66c", Color: "#9b59b6"},
	"java":         {Icon: "\ue738", Color: "#e67e22"},
	"jenkinsfile":  {Icon: "\ue767", Color: "#e74c3c"},
	"jl":           {Icon: "\ue624", Color: "#2ecc71"},
	"js":           {Icon: "\ue781", Color: "#f39c12"},
	"json":         {Icon: "\ue60b", Color: "#f1c40f"},
	"jsx":          {Icon: "\ue7ba", Color: "#e67e22"},
	"key":          {Icon: "\uf43d", Color: "#f1c40f"},
	"ko":           {Icon: "\uebc6", Color: "#9b59b6"},
	"kt":           {Icon: "\ue634", Color: "#2980b9"},
	"less":         {Icon: "\ue758", Color: "#3498db"},
	"lock":         {Icon: "\uf023", Color: "#f1c40f"},
	"log":          {Icon: "\uf18d", Color: "#7f8c8d"},
	"lua":          {Icon: "\ue620", Color: "#e74c3c"},
	"maintainers":  {Icon: "\uf0c0", Color: "#7f8c8d"},
	"makefile":     {Icon: "\ue20f", Color: "#3498db"},
	"md":           {Icon: "\uf48a", Color: "#7f8c8d"},
	"mjs":          {Icon: "\ue718", Color: "#f39c12"},
	"ml":           {Icon: "󰘧", Color: "#2ecc71"},
	"mustache":     {Icon: "\ue60f", Color: "#e67e22"},
	"nc":           {Icon: "󰋁", Color: "#f1c40"},
	"nim":          {Icon: "\ue677", Color: "#3498db"},
	"nix":          {Icon: "\uf313", Color: "#f39c12"},
	"npmignore":    {Icon: "\ue71e", Color: "#e74c3c"},
	"package":      {Icon: "󰏗", Color: "#9b59b6"},
	"passwd":       {Icon: "\uf023", Color: "#f1c40f"},
	"patch":        {Icon: "\uf440", Color: "#e67e22"},
	"pdf":          {Icon: "\uf1c1", Color: "#d35400"},
	"php":          {Icon: "\ue608", Color: "#9b59b6"},
	"pl":           {Icon: "\ue7a1", Color: "#3498db"},
	"prisma":       {Icon: "\ue684", Color: "#9b59b6"},
	"ppt":          {Icon: "\uf1c4", Color: "#c0392b"},
	"psd":          {Icon: "\ue7b8", Color: "#3498db"},
	"py":           {Icon: "\ue606", Color: "#3498db"},
	"r":            {Icon: "\ue68a", Color: "#9b59b6"},
	"rb":           {Icon: "\ue21e", Color: "#9b59b6"},
	"rdb":          {Icon: "\ue76d", Color: "#9b59b6"},
	"rpm":          {Icon: "\uf17c", Color: "#d35400"},
	"rs":           {Icon: "\ue7a8", Color: "#f39c12"},
	"rss":          {Icon: "\uf09e", Color: "#c0392b"},
	"rst":          {Icon: "󰅫", Color: "#2ecc71"},
	"rubydoc":      {Icon: "\ue73b", Color: "#e67e22"},
	"sass":         {Icon: "\ue603", Color: "#e74c3c"},
	"scala":        {Icon: "\ue737", Color: "#e67e22"},
	"shell":        {Icon: "\uf489", Color: "#2ecc71"},
	"shp":          {Icon: "󰙞", Color: "#f1c40f"},
	"sol":          {Icon: "󰡪", Color: "#3498db"},
	"sqlite":       {Icon: "\ue7c4", Color: "#27ae60"},
	"styl":         {Icon: "\ue600", Color: "#e74c3c"},
	"svelte":       {Icon: "\ue697", Color: "#ff3e00"},
	"swift":        {Icon: "\ue755", Color: "#ff6f61"},
	"tex":          {Icon: "\u222b", Color: "#9b59b6"},
	"tf":           {Icon: "\ue69a", Color: "#2ecc71"},
	"toml":         {Icon: "󰅪", Color: "#f39c12"},
	"ts":           {Icon: "󰛦", Color: "#2980b9"},
	"twig":         {Icon: "\ue61c", Color: "#9b59b6"},
	"txt":          {Icon: "\uf15c", Color: "#7f8c8d"},
	"vagrantfile":  {Icon: "\ue21e", Color: "#3498db"},
	"video":        {Icon: "\uf03d", Color: "#c0392b"},
	"vim":          {Icon: "\ue62b", Color: "#019833"},
	"vue":          {Icon: "\ue6a0", Color: "#41b883"},
	"windows":      {Icon: "\uf17a", Color: "#4a90e2"},
	"xls":          {Icon: "\uf1c3", Color: "#27ae60"},
	"xml":          {Icon: "\ue796", Color: "#3498db"},
	"yml":          {Icon: "\ue601", Color: "#f39c12"},
	"zig":          {Icon: "\ue6a9", Color: "#9b59b6"},
	"zip":          {Icon: "\uf410", Color: "#e74c3c"},
}

var aliases = map[string]string{
	"dart":             "dart",
	"apk":              "android",
	"gradle":           "android",
	"ds_store":         "apple",
	"localized":        "apple",
	"m":                "apple",
	"mm":               "apple",
	"s":                "asm",
	"aac":              "audio",
	"alac":             "audio",
	"flac":             "audio",
	"m4a":              "audio",
	"mka":              "audio",
	"mp3":              "audio",
	"ogg":              "audio",
	"opus":             "audio",
	"wav":              "audio",
	"wma":              "audio",
	"bson":             "binary",
	"feather":          "binary",
	"mat":              "binary",
	"o":                "binary",
	"pb":               "binary",
	"pickle":           "binary",
	"pkl":              "binary",
	"tfrecord":         "binary",
	"conf":             "cfg",
	"config":           "cfg",
	"cljc":             "clj",
	"cljs":             "clj",
	"editorconfig":     "conf",
	"rc":               "conf",
	"c++":              "cpp",
	"cc":               "cpp",
	"cxx":              "cpp",
	"scss":             "css",
	"sql":              "db",
	"docx":             "doc",
	"gdoc":             "doc",
	"dockerignore":     "dockerfile",
	"epub":             "ebook",
	"ipynb":            "ebook",
	"mobi":             "ebook",
	"env":              "env",
	".env.local":       "env",
	"local":            "env",
	"f03":              "f",
	"f77":              "f",
	"f90":              "f",
	"f95":              "f",
	"for":              "f",
	"fpp":              "f",
	"ftn":              "f",
	"eot":              "font",
	"otf":              "font",
	"ttf":              "font",
	"woff":             "font",
	"woff2":            "font",
	"fsi":              "fs",
	"fsscript":         "fs",
	"fsx":              "fs",
	"dna":              "gb",
	"gitattributes":    "git",
	"gitconfig":        "git",
	"gitignore":        "git",
	"gitignore_global": "git",
	"gitmirrorall":     "git",
	"gitmodules":       "git",
	"gltf":             "glp",
	"gsh":              "groovy",
	"gvy":              "groovy",
	"gy":               "groovy",
	"h++":              "h",
	"hh":               "h",
	"hpp":              "h",
	"hxx":              "h",
	"lhs":              "hs",
	"htm":              "html",
	"xhtml":            "html",
	"bmp":              "image",
	"cbr":              "image",
	"cbz":              "image",
	"dvi":              "image",
	"eps":              "image",
	"gif":              "image",
	"ico":              "image",
	"jpeg":             "image",
	"jpg":              "image",
	"nef":              "image",
	"orf":              "image",
	"pbm":              "image",
	"pgm":              "image",
	"png":              "image",
	"pnm":              "image",
	"ppm":              "image",
	"pxm":              "image",
	"sixel":            "image",
	"stl":              "image",
	"svg":              "image",
	"tif":              "image",
	"tiff":             "image",
	"webp":             "image",
	"xpm":              "image",
	"disk":             "iso",
	"dmg":              "iso",
	"img":              "iso",
	"ipsw":             "iso",
	"smi":              "iso",
	"vhd":              "iso",
	"vhdx":             "iso",
	"vmdk":             "iso",
	"jar":              "java",
	"cjs":              "js",
	"properties":       "json",
	"webmanifest":      "json",
	"tsx":              "jsx",
	"cjsx":             "jsx",
	"cer":              "key",
	"crt":              "key",
	"der":              "key",
	"gpg":              "key",
	"p7b":              "key",
	"pem":              "key",
	"pfx":              "key",
	"pgp":              "key",
	"license":          "key",
	"codeowners":       "maintainers",
	"credits":          "maintainers",
	"cmake":            "makefile",
	"justfile":         "makefile",
	"markdown":         "md",
	"mkd":              "md",
	"rdoc":             "md",
	"readme":           "md",
	"mli":              "ml",
	"sml":              "ml",
	"netcdf":           "nc",
	"brewfile":         "package",
	"cargo.toml":       "package",
	"cargo.lock":       "package",
	"go.mod":           "package",
	"go.sum":           "package",
	"pyproject.toml":   "package",
	"poetry.lock":      "package",
	"package.json":     "package",
	"pipfile":          "package",
	"pipfile.lock":     "package",
	"php3":             "php",
	"php4":             "php",
	"php5":             "php",
	"phpt":             "php",
	"phtml":            "php",
	"gslides":          "ppt",
	"pptx":             "ppt",
	"pxd":              "py",
	"pyc":              "py",
	"pyx":              "py",
	"whl":              "py",
	"rdata":            "r",
	"rds":              "r",
	"rmd":              "r",
	"gemfile":          "rb",
	"gemspec":          "rb",
	"guardfile":        "rb",
	"procfile":         "rb",
	"rakefile":         "rb",
	"rspec":            "rb",
	"rspec_parallel":   "rb",
	"rspec_status":     "rb",
	"ru":               "rb",
	"erb":              "rubydoc",
	"slim":             "rubydoc",
	"awk":              "shell",
	"bash":             "shell",
	"bash_history":     "shell",
	"bash_profile":     "shell",
	"bashrc":           "shell",
	"csh":              "shell",
	"fish":             "shell",
	"ksh":              "shell",
	"sh":               "shell",
	"zsh":              "shell",
	"zsh-theme":        "shell",
	"zshrc":            "shell",
	"plpgsql":          "sql",
	"plsql":            "sql",
	"psql":             "sql",
	"tsql":             "sql",
	"sl3":              "sqlite",
	"sqlite3":          "sqlite",
	"stylus":           "styl",
	"cls":              "tex",
	"avi":              "video",
	"flv":              "video",
	"m2v":              "video",
	"mkv":              "video",
	"mov":              "video",
	"mp4":              "video",
	"mpeg":             "video",
	"mpg":              "video",
	"ogm":              "video",
	"ogv":              "video",
	"vob":              "video",
	"webm":             "video",
	"vimrc":            "vim",
	"bat":              "windows",
	"cmd":              "windows",
	"exe":              "windows",
	"csv":              "xls",
	"gsheet":           "xls",
	"xlsx":             "xls",
	"plist":            "xml",
	"xul":              "xml",
	"yaml":             "yml",
	"7z":               "zip",
	"Z":                "zip",
	"bz2":              "zip",
	"gz":               "zip",
	"lzma":             "zip",
	"par":              "zip",
	"rar":              "zip",
	"tar":              "zip",
	"tc":               "zip",
	"tgz":              "zip",
	"txz":              "zip",
	"xz":               "zip",
	"z":                "zip",
}

var folders = map[string]IconStyle{
	".atom":                 {Icon: "\ue764", Color: "#66595c"}, // Atom folder - Dark gray
	".aws":                  {Icon: "\ue7ad", Color: "#ff9900"}, // AWS folder - Orange
	".docker":               {Icon: "\ue7b0", Color: "#0db7ed"}, // Docker folder - Blue
	".gem":                  {Icon: "\ue21e", Color: "#e9573f"}, // Gem folder - Red
	".git":                  {Icon: "\ue5fb", Color: "#f14e32"}, // Git folder - Red
	".git-credential-cache": {Icon: "\ue5fb", Color: "#f14e32"}, // Git credential cache folder - Red
	".github":               {Icon: "\ue5fd", Color: "#000000"}, // GitHub folder - Black
	".npm":                  {Icon: "\ue5fa", Color: "#cb3837"}, // npm folder - Red
	".nvm":                  {Icon: "\ue718", Color: "#cb3837"}, // nvm folder - Red
	".rvm":                  {Icon: "\ue21e", Color: "#e9573f"}, // rvm folder - Red
	".Trash":                {Icon: "\uf1f8", Color: "#7f8c8d"}, // Trash folder - Light gray
	".vscode":               {Icon: "\ue70c", Color: "#007acc"}, // VSCode folder - Blue
	".vim":                  {Icon: "\ue62b", Color: "#019833"}, // Vim folder - Green
	"config":                {Icon: "\ue5fc", Color: "#ffb86c"}, // Config folder - Light orange
	"folder":                {Icon: "", Color: "NONE"},         // Generic folder - Dark yellowish
	"hidden":                {Icon: "\uf023", Color: "#75715e"}, // Hidden folder - Dark yellowish
	"node_modules":          {Icon: "\ue5fa", Color: "#cb3837"}, // Node modules folder - Red
}
