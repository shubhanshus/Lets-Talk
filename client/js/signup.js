$(document).ready(function(){

    $('#signup').click(function() {
    	var uname=$("#username").val();
		var password1=$("#password1").val(); 
		var password2=$("#password2").val(); 
		var email=$("#email").val(); 

		//check required options
		if(uname=='' || password1=='' || password2=='' || email==''){
			alert("Please input all the information, thank you");
			return;
		}


		if(password1 != password2){
			alert("2 passwords are not the same, please check");
			return;
		}
	
	});

});