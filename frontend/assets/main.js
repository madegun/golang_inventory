

document.getElementById("query_form").addEventListener("submit", event => {
	getAjax("/api/device/" +  document.getElementById("device_query").value, function(data) {
		document.body.innerHTML += data;
		console.log(data)
	})
	event.preventDefault();
})