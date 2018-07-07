function replace(divId, contents, contents2){
    var div = document.getElementById(divId);
    if(div.innerHTML == contents){
        div.innerHTML = contents2
    }else{
        div.innerHTML = contents
    }
}