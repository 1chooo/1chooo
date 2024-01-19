# Code Style About Hugo


```go
package Hugo

func aboutMe() {
    type Me struct {
        Pronouns string
        Code     string
        Job      string
        Company  string
        Location string
    }

    var me = Me{
        Pronouns: "He/Him",
        Code:     "Go, Python, C++, C, Java, JavaScript, HTML, CSS",
        Job:      "Software Engineer",
        Company:  "Pegatron Corp.",
        Location: "Taiwan",
    }
    _ = me
}

func getKnowledge() {
    type Knowledge struct {
        Languages []string
        Tools     []string
        Hobbies   []string
    }

    var knowledge = Knowledge{
        Languages: []string{"Go", "Python", "C++", "C", "Java", "JavaScript", "HTML", "CSS"},
        Tools:     []string{"Linux", "Git", "Docker", "AWS", "GCP", "Kubernetes", "CI/CD"},
        Hobbies:   []string{"Gaming", "Reading", "Music", "Movies"},
    }
    _ = knowledge
}

func getFutureGoal() {
    type FutureGoal struct {
        Life string
        Work string
    }

    var futureGoal = FutureGoal{
        Life: "Love what my love",
        Work: "Do what I want",
    }
    _ = futureGoal
}
```


<!-- ```ts
public sealed class About : Me {
    private const string MY_DREAM = "To be a Fullstack Developer.";
    private const string MY_QUOTE = "It is no use doing what you like; you have got to like what you do."

    public object GetCurrentWorkplace() {
        return new {
            Workplace = new {
                Company = "Thoughtworks",
                Position = "China"
            }
        };
    }

    public IEnumerable<string> GetDailyKnowledge() {
        return new[] {
            "C#", "ASP.NET Core", "WPF", "UWP",
            "React.js", "Vue.js", "HTML", "JavaScript", "CSS",
            "SQL Server", "Oracle", "MySQL",
            "Nginx", "IIS", "Powershell", "Python",
            "Windows", "Ubuntu", "MacOS",
            "Azure Devops", "Github Action", "CI&CD",
            "TDD", "DDD"
        };
    }

    public IDictionary<string, string> GetFutureGoal() {
        return new Dictionary<string, string>() {
            { "Life", "Love what my love" },
            { "Work", "Do what I want" }
        };
    }
}
``` -->