"use strict";
// add event handlers
$("#sendValues").on("click",function(){
  var values ={};
  values.startPopulation = $("#startPop").val();
  values.growthRate = $("#growthRate").val();
  //var dataString = JSON.stringify(values);
  var dataString = values;
  $.ajax(
    "http://localhost:8080/population",
    {
      "type":"POST",
      "dataType":"json",
      "data": dataString,
      "complete":function( jqXHR ,  textStatus){
        $("#response").html($("#response").html() + "\n" + textStatus)
      },
      "success":function(){
        $("#response").html("values submitted successfully")
      },
      "error":function( jqXHR ,  textStatus,  errorThrown){
        $("#response").html("error! " + errorThrown)
      }
    }
  );
});
$("#getValues").on("click",function(){
  var values ={};
  values.startPopulation = $("#startPop").val();
  values.growthRate = $("#growthRate").val();
  var dataString = JSON.stringify(values);
  $.ajax(
    "http://localhost:8080/population",
    {
      "type":"GET",
      "dataType":"json",
      "success":function(data, textStatus,jqXHR){
        if(typeof(data) !== "undefined"){
          if(typeof(data.CurrentPopulation) !== "undefined"){
            $("#currentPop").html(data.CurrentPopulation);
            return;
          }
        }
        $("#response").html("GET success :)");
      },
      "error":function( jqXHR,  textStatus,  errorThrown){
        $("#response").html("GET error! " + errorThrown);
      },
      "complete":function( jqXHR,  textStatus){
        $("#response").html($("#response").html() + "\n" + textStatus);
      }
    }
  );
});