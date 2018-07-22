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
        <div class="content">
            <h1 style="color: white">Welcome, {{.Username}}</h1>
        </div>
        <script src="/assets/js/scripts.js"></script>
    </body>
</html>