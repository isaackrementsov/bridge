{{define "content"}}
    <div class="content">
        <div class="form inline">
            <div class="loginquestion child">
                <img src="/assets/img/bridge2.svg" class="formimg">
                <h2>Already have an account?</h2>
                <a class="button" href="/login">Log in</a>
            </div>
            <div class="child signup">
                <h1>Sign up</h1>
                <form method="post">
                    <input type="text" name="username" placeholder="Username" class="input big">
                    <input type="email" name="email" placeholder="Email" class="input big">
                    <input type="password" name="password" placeholder="Password" class="input big">
                    <input type="submit" value="Sign up" class="button ctainverse big">
                </form>
            </div>
        </div>
    </div>
{{end}}