package socialmedia

type SocialMedia interface {
	Feed() []string
	Fame() int
}

type Facebook struct {
	UserName string
	Email    string
}

func (f Facebook) Feed() []string {
	return []string{
		"Hello, my first post",
		"This is getting really awesome",
		"I had my first picnic in Germany",
	}
}

func (f Facebook) Fame() int {
	return 208
}
