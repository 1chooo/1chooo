// Hugo.go

package Hugo

type Me struct {
	Name      string
	Pronouns  string
	Skill     string
	FrameWork string
	DevOps    string
	Job       string
	Company   string
	Location  string
}

func aboutMe() Me {
	return Me {
		Name:      "Hugo",
		Pronouns:  "He/Him",
		Skill:     "Go, Python, TypeScript, C++, C, Java, JavaScript",
		FrameWork: "React, FastAPI, Flask",
		DevOps:    "GitHub Actions, Docker, Kubernetes, AWS, Azure",
		Job:       "Cloud Developer",
		Company:   "eCloudvalley Technology Co., Ltd.",
		Location:  "Taiwan",
	}
}

type Knowledge struct {
	Languages []string
	Tools     []string
	Hobbies   []string
}

func getKnowledge() Knowledge {

	return Knowledge {
		Languages: []string{"Go", "Python", "C++", "C", "Java", "JavaScript", "HTML", "CSS"},
		Tools:     []string{"Linux", "Git", "Docker", "AWS", "GCP", "Kubernetes", "CI/CD"},
		Hobbies:   []string{"Gaming", "Reading", "Music", "Movies"},
	}
}

type FutureGoal struct {
	Life string
	Work string
}

func getFutureGoal() FutureGoal {

	return FutureGoal {
		Life: "Love what my love",
		Work: "Do what I want",
	}
}
