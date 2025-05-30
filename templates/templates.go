package templates

import (
	"embed"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

type (
	Layout string
	Page   string
)

const (
	LayoutMain Layout = "main"
	LayoutAuth Layout = "auth"
	LayoutHTMX Layout = "htmx"
)

const (
	PageAbout          Page = "about"
	PageCache          Page = "cache"
	PageContact        Page = "contact"
	PageError          Page = "error"
	PageForgotPassword Page = "forgot-password"
	PageHome           Page = "home"
	PageLogin          Page = "login"
	PageRegister       Page = "register"
	PageResetPassword  Page = "reset-password"
	PageSearch         Page = "search"
	PageTask           Page = "task"
	PageClickToEdit    Page = "click-to-edit"
	PageDashboard      Page = "dashboard"
	PageStripe         Page = "stripe"
	PageStripeSuccess  Page = "stripe-success"
	PageStripeCancel   Page = "stripe-cancel"
	// -------- Model Pages --------
	PageModelPrice         Page = "model-price"
	PageModelPriceList     Page = "model-price-list"
	PageModelPriceForm     Page = "model-price-form"
	PageModelProduct       Page = "model-product"
	PageModelProductList   Page = "model-product-list"
	PageModelProductForm   Page = "model-product-form"
	PageModelModelName     Page = "model-modelname"
	PageModelModelNameList Page = "model-modelname-list"
	PageModelModelNameForm Page = "model-modelname-form"
	PageAITask             Page = "ai-task"
)

//go:embed *
var templates embed.FS

// Get returns a file system containing all templates via embed.FS
func Get() embed.FS {
	return templates
}

// GetOS returns a file system containing all templates which will load the files directly from the operating system.
// This should only be used for local development in order to facilitate live reloading.
func GetOS() fs.FS {
	// Gets the complete templates directory path
	// This is needed in case this is called from a package outside of main, such as within tests
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	p := filepath.Join(filepath.Dir(d), "templates")
	return os.DirFS(p)
}
