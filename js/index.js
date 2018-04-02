$(document).ready(function(){

 	var uname=$("#uname").text();
	$("#logout").hide();
	$("#cancel").hide();

	if(uname != ''){
			$("#signup").hide();
			$("#login").hide();
			$("#logout").show();
			$("#cancel").show();
	}


    
});