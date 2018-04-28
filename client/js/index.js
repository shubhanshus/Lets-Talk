$(function(){

 	var uname=$("#uname").text();
	$("#logout").hide();
	$("#cancel").hide();
	$("#follow").hide();
	$("#unfollow").hide();

	if(uname != ''){
			$("#signup").hide();
			$("#login").hide();
			$("#logout").show();
			$("#cancel").show();
			$("#follow").show();
			$("#unfollow").show();
	}

	$( "#talk" ).click(function() {
		if(uname == ''){
			alert("Please login or sign up first, thanks");
			window.history.back();
		}
  		
	});

	$.ajax({
        type: "GET",
		url: document.location.href+'/list',
		dataType:'json',
		contentType:"application/json",
		success:function(data){
			console.log(data);
			var talklist;
			var count = 0;
			if(data == null){
				talklist = "<p>There is no post.</p>";
			}else{
				talklist = "";

				for (var i = data.length - 1; i >= 0; i--) {
					if(data[i].email == ""){
						data[i].email = "Anonymous";
					}
					if($.cookie("Count"+i) == undefined){
						count = 0;
					}else{
						count = $.cookie("Count"+i);
					}
					
					talklist += "<div class='talklist'><div class='profile'><img src='/img/profile.jpg' height ='50' width='50' /></div>"+
								"<div class='username'><h>"+data[i].email+"</h></div>"+
								"<div class='content'><p>"+data[i].talk+"</p></div>"+
								"<div class='pdate'><p>"+data[i].date+"</p></div>"+
								"<div id='heart'><a id='"+i+"' href=''><img src='/img/heart.png' height ='20' width='30' /></a></div><div id='num"+i+"' class='num'><p>"+count+"</p></div></div>"
					
					console.log(data[i].email, data[i].talk, data[i].date, i);
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

        //set expire time 20 mins
        var exp = new Date();
		var time = exp.getTime();
		time += 20 * 60 * 1000;
		exp.setTime(time);
		document.cookie="Count"+likeId+"="+count+";expires="+exp.toGMTString()+";path=/";
});  

