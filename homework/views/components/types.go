package components

type LayoutProps struct {
	Title           string
	MetaDescription string
}

type NavItemProps struct {
	Name  string
	Link  string
	Image string
}

type TopicCardProps struct {
	Title string
	Text  string
	Img   string
}

type BlogCardProps struct {
	Author        string
	AuthorImg     string
	ArticleHeader string
	Article       string
	Date          string
	BlogImg       string
}

type ArrowDirection string

const (
	ArrowNone  = ""
	ArrowLeft  = "chevron-left"
	ArrowRight = "chevron-right"
)

type ButtonVariant string

const (
	ButtonPrimary   ButtonVariant = "primary"
	ButtonSecondary ButtonVariant = "secondary"
	ButtonIconOnly  ButtonVariant = "icon-only"
)

type ButtonProps struct {
	Text    string
	Arrow   ArrowDirection
	Variant ButtonVariant
	Link    string
}

type InputType string

const (
	InputText     InputType = "text"
	InputEmail    InputType = "email"
	InputPassword InputType = "password"
)

type InputProps struct {
	Label       string
	Name        string
	Type        InputType
	Placeholder string
	Required    bool
	Error       string
}
