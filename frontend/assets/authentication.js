function postAjax(url, data, success) {
	var params = typeof data == 'string' ? data : Object.keys(data).map(function(k) {
		return encodeURIComponent(k) + '=' + encodeURIComponent(data[k])
	}).join('&');
	var xhr = window.XMLHttpRequest ? new XMLHttpRequest() : new ActiveXObject("Microsoft.XMLHTTP");
	xhr.open('POST', url);
	xhr.onreadystatechange = function() {
		if (xhr.readyState > 3 && xhr.status == 200) {
			success(xhr.responseText);
		}
	};
	xhr.setRequestHeader('X-Requested-With', 'XMLHttpRequest');
	xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
	xhr.send(params);
	return xhr;
}

function setCookie(cname, cvalue, exdays) {
	var d = new Date();
	d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
	var expires = "expires=" + d.toUTCString();
	document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/" + ";Secure";
}

function getCookie(cname) {
	var name = cname + "=";
	var ca = document.cookie.split(';');
	for (var i = 0; i < ca.length; i++) {
		var c = ca[i];
		while (c.charAt(0) == ' ') {
			c = c.substring(1);
		}
		if (c.indexOf(name) == 0) {
			return c.substring(name.length, c.length);
		}
	}
	return "";
}
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