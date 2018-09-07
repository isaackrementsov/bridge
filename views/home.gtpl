{{define "content"}}
    <div class="content profile">
        <h1>Welcome, {{.Session.Username}}</h1>
        {{if .DB}}
            <h2 class="subtitle">Your domains</h2>
            {{range .DB}}
                {{.Name}}
            {{end}}
        {{else}}
            <h2 class="subtitle light">You have no domains yet</h2>
        {{end}}
        <div style="display: block"><a class="button cta ibl" onclick="opendiv('newDomain')"><i class="material-icons">add</i></a></div>
    </div>
{{end}}