<!DOCTYPE html>
<html>
    <head>
        <title>Bridge</title>
        <link rel="stylesheet" type="text/css" href="/assets/css/main.css">
        <link href="https://fonts.googleapis.com/icon?family=Material+Icons"rel="stylesheet">
    </head>
    <body>
        <ul class="navbar">
            <li><img src="/assets/img/bridge.svg"></li>
            <li><a href="/signup" class="hover">Sign up</a></li>  
            <li><a href="/login">Log in</a></li>
        </ul>
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
        <script src="/assets/js/scripts.js"></script>
    </body>
</html>