function setPosition(position){
  theUrl = "/position/"+position;
  var callback = function(responseText){
    console.log(responseText);
  }
  httpGetAsync(theUrl, callback);
}

function httpGetAsync(theUrl, callback)
{
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.onreadystatechange = function() {
        if (xmlHttp.readyState == 4 && xmlHttp.status == 200)
            callback(xmlHttp.responseText);
    }
    xmlHttp.open("GET", theUrl, true); // true for asynchronous
    xmlHttp.send(null);
}

function onLoadBody(){
    var  tempWidth = document.documentElement.clientWidth;
    var tempHeight = document.documentElement.clientHeight;
    var elements = document.getElementsByClassName("bttnPosition");
    for (var i=0; i<elements.length; i++){
        elements[i].style.height=tempHeight/3-10+"px";
        //elements[i].style.width=tempWidth/3-5+"px";
    }
}
