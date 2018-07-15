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
        <div class="main inline">
            <div class="child img">
                <img src="/assets/img/bridge2.svg">
            </div>
            <div class="child text">
                <h1>Welcome to Bridge</h1>
                <p>
                    Bridge is a collection of services to use backend functionality on a frontend website. 
                    Whether you need database access or session storage, Bridge provides a simple 
                    JavaScript API to implement advanced server-side features.
                </p>
                <div class="inline links">
                    <a href="#features"class="button">Features</a>
                    <a href="/signup"class="button cta">Get started</a>
                </div>
            </div>
        </div>
        <div class="features">
            <a href="#features">
                <i class="material-icons"id="features" style="font-size:6vw">arrow_drop_down</i>
            </a>
            <h2>Features</h2>
            <div class="square" style="background-color:#90CAF9" id="db">
                <h2><i class="material-icons">storage</i></h2>
                <h3>Database</h3>
            </div>
            <div class="square" style="background-color:#64B5F6" id="sess">
                <h2><i class="material-icons">restore</i></h2>
                <h3>Sessions</h3>
            </div>
            <div class="square" style="background-color:#42A5F5" id="up">
                <h2><i class="material-icons">cloud_upload</i></h2>
                <h3>File upload</h3>
            </div>
            <div class="square" style="background-color:#2196F3" id="email">
                <h2><i class="material-icons">mail</i></h2>
                <h3>Email</h3>
            </div>
            <p style="width:50vw">Whatever you need, Bridge has it. 
                With a continuously expanding selection of backend tools surrounding a 
                core feature set, you can simply plug scripts in and play. Pricing is done per plugin, so you can purchase as few or as many as you desire.</p>
            <a class="normal" href="https://isaackrementsov.github.io/bridge">Read the docs</a>
            <a href="/signup" class="button ctainverse">Let's go!</a>
        </div>
        <script src="/assets/js/scripts.js"></script>
    </body>
</html>