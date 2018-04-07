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
			window.history.back();
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
			var count = 0;
			if(data == null){
				talklist = "<p>There is not post.</p>";
			}else{
				talklist = "";

				for (var i = data.length - 1; i >= 0; i--) {
					if(data[i].UserName == ""){
						data[i].UserName = "Anonymous";
					}
					if($.cookie("Count"+i) == undefined){
						count = 0;
					}else{
						count = $.cookie("Count"+i);
					}
					
					talklist += "<div class='talklist'><div class='profile'><img src='/img/profile.jpg' height ='50' width='50' /></div>"+
								"<div class='username'><h>"+data[i].UserName+"</h></div>"+
								"<div class='content'><p>"+data[i].Talk+"</p></div>"+
								"<div class='pdate'><p>"+data[i].Date+"</p></div>"+
								"<div id='heart'><a id='"+i+"' href=''><img src='/img/heart.png' height ='20' width='30' /></a></div><div id='num"+i+"' class='num'><p>"+count+"</p></div></div>"
					
					console.log(data[i].UserName, data[i].Talk, data[i].Date, i);
				}
			}
			
			$('.talklists').html(talklist);
   
        },
		error:function(){
			alert("connection error");
		}
    });  

   

});

$(document).on('click', '#heart a', function (){
        likeId=this.id;
        //alert(likeId);
        var count = $("#num"+likeId).text();
        count++;
        //$("#num"+likeId).val() = count;
        alert("Liked");
		document.cookie="Count"+likeId+"="+count;
});  

