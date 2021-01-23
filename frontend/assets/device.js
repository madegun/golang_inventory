var device_id = window.location.pathname.split("/device/")[1];

window.onload = function() {

    getAjax("/api/device/" + device_id, function(data) {
		parsed_data = JSON.parse(data);
		if (parsed_data.error == 1) {
			var error_alert = document.createElement("div")
			error_alert.className = "alert alert-danger"
			error_alert.innerHTML = parsed_data.message
			document.body.append(error_alert)			
		} else {
            for (var key in parsed_data) {
                if(parsed_data.hasOwnProperty(key)) {
                    if(key != "delivery" && key != "donator" && key != "devid" && key != "ID" && key != "CreatedAt" && key !="UpdatedAt" && key != "DeletedAt") {
                        Group = document.createElement("div")
                        Group.className = "input-group mb-3"
                        Label = document.createElement("span")
                        Label.className = "input-group-text"
                        Label.innerHTML = EN[key]
                        Input = document.createElement("input")
                        Input.type = "text"
                        Input.className = "form-control"
                        Input.value = parsed_data[key]
                        Group.append(Label)
                        Group.append(Input)
                        document.getElementById("device").append(Group)
                    }
                }
            }
            for(var key in parsed_data["delivery"]) {
                if (parsed_data["delivery"].hasOwnProperty(key)) {
                    if(key != "DeviceID" && key != "ID" && key != "CreatedAt" && key !="UpdatedAt" && key != "DeletedAt") {
                        Group = document.createElement("div")
                        Group.className = "input-group mb-3"
                        Label = document.createElement("span")
                        Label.className = "input-group-text"
                        Label.innerHTML = EN[key]
                        Input = document.createElement("input")
                        Input.type = "text"
                        Input.className = "form-control"
                        Input.value = parsed_data["delivery"][key]
                        Group.append(Label)
                        Group.append(Input)
                        document.getElementById("delivery").append(Group)
                    }
                }
            }
            for(var key in parsed_data["donator"]) {
                if (parsed_data["donator"].hasOwnProperty(key)) {
                    if(key != "DeviceID" && key != "ID" && key != "CreatedAt" && key !="UpdatedAt" && key != "DeletedAt") {
                        Group = document.createElement("div")
                        Group.className = "input-group mb-3"
                        Label = document.createElement("span")
                        Label.className = "input-group-text"
                        Label.innerHTML = EN[key]
                        Input = document.createElement("input")
                        Input.type = "text"
                        Input.className = "form-control"
                        Input.value = parsed_data["donator"][key]
                        Group.append(Label)
                        Group.append(Input)
                        document.getElementById("donator").append(Group)
                    }
                }
            }
		}
	})
}