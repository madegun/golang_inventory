
document.getElementById("login_form").addEventListener("submit", event => {
	postAjax("/api/login", {
		"email": document.getElementById("email").value,
		"password": document.getElementById("password").value
	}, function(data) {
		parsed_data = JSON.parse(data);
		if (parsed_data.error == 1) {
			if (document.getElementById("login_error")) {
				document.getElementById("login_error").innerHTML = parsed_data.message
			} else {
				var error_alert = document.createElement("div")
				error_alert.id = "login_error"
				error_alert.className = "alert alert-danger"
				error_alert.innerHTML = parsed_data.message
				document.body.append(error_alert)
			}
		} else {
			
			setCookie("token", parsed_data.token, 7)
			window.location.replace("/dashboard")
		}
	})
	event.preventDefault();
})