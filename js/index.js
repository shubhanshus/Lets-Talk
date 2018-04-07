$(function(){

 	var uname=$("#uname").text();
	$("#logout").hide();
	$("#cancel").hide();

	if(uname != ''){
			$("#signup").hide();
			$("#login").hide();
			$("#logout").show();
			$("#cancel").show();
	}

	$( "#talk" ).click(function() {
		if(uname == ''){
			alert("Please login or sign up first, thanks");
			return;
		}
  		
	});

	$.ajax({
        type: "GET",
		url: 'http://localhost:8080/list',
		dataType:'json',
		contentType:"application/json",
		success:function(data){
			console.log(data);
			var talklist;

			if(data == "undefined"){
				talklist = "<p>There is not post.</p>";
			}else{
				talklist = "";

				for (var i = data.length - 1; i >= 0; i--) {
					if(data[i].UserName == ""){
						data[i].UserName = "Anonymous";
					}
					talklist += "<div class='talklist'><div class='profile'><img src='/img/profile.jpg' height ='50' width='50' /></div>"+
								"<div class='username'><h>"+data[i].UserName+"</h></div>"+
								"<div class='content'><p>"+data[i].Talk+"</p></div>"+
								"<div class='pdate'><p>"+data[i].Date+"</p></div>"+
								"<div class='heart'><a href=''><img src='/img/heart.png' height ='20' width='30' /></a></div><div id='like'><p>Like</p></div></div>"
					
					console.log(data[i].UserName, data[i].Talk, data[i].Date);
				}
			}
			
			$('.talklists').html(talklist);
   
        },
		error:function(){
			alert("connection error");
		}
    });


    
});