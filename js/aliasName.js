name = (typeof name == "undefined"? "undefined" : name)
output_tmp = {}
if(name != "undefined"){
  input[0].Counter = name
  output_tmp = _.first(input, 1)
}else{
  output_tmp = input
}
output = JSON.stringify(output_tmp)
