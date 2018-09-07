{{define "content"}}
    <div class="content">
        <div class="form inline">
            <div class="loginquestion child">
                <img src="/assets/img/bridge2.svg" class="formimg">
                <h2>Don't have an account?</h2>
                <a class="button" href="/signup">Sign up</a>
            </div>
            <div class="child signup">
                <h1>Log in</h1>
                <form method="post">
                    <input type="text" name="username" placeholder="Username" class="input big">
                    <input type="password" name="password" placeholder="Password" class="input big">
                    <input type="submit" value="Log in" class="button ctainverse big">
                </form>
            </div>
        </div>
    </div>
{{end}}