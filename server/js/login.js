$(document).ready(function(){

    $('#login').click(function() {

		var uname=$("#username").val();
		var password=$("#password").val(); 
		//check required options
		if(uname=='' || password==''){
			alert("Please input all the information, thank you");
			return;
		}
	
	});

});