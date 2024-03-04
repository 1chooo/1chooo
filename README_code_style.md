# Code Style About Hugo

```go
package Hugo

type Me struct {
	Name      string
	Pronouns  string
	Contact   string
	Job       string
	Company   string
	Location  string
	Hobbies   string
}

func aboutMe() Me {

	return Me{
		Name:      "Hugo",
		Pronouns:  "He/Him",
		Contact:   "hugo970217@gmail.com",
		Job:       "Cloud Developer",
		Company:   "eCloudValley Technology Co., Ltd.",
		Location:  "Taiwan 🇹🇼",
		Hobbies:   "Reading 📖, Music 🎶, Movies 🍿, Photography 📸",
	}
}

type Skill struct {
	Languages  []string
	FrameWork  []string
	DevOps     []string
}

func getSkill() Skill {

	return Skill{
		Languages: []string{"Go", "Python", "TypeScript", "C++", "C", "Java", "JavaScript"},
		FrameWork: []string{"React", "FastAPI", "Flask", "Gin", "Langchain"},
		DevOps:    []string{"AWS", "Linux", "GitHub Actions", "Docker", "Kubernetes", "Azure"},
	}
}

type Working struct {
	Company   []string
}

type FutureGoal struct {
	Life       string
	Work       string
	Education  string
}

func getFutureGoal() FutureGoal {

	return FutureGoal{
		Life: "Love what my love",
		Work: "Do what I want",
		Education: "Study in the field of Computer Science in the United States 🇺🇸",
	}
}

func getQuote() string {
	return "Dreams come TRUE, when 🫵🏻 don't SLEEP...💤"
}
```
