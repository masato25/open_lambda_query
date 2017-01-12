# 製作一個Lambda function
參考網址:​
* [otto](https://github.com/robertkrimen/otto)

* ​[underscorejs](http://underscorejs.org/)

## 以製作top function 為例:

top function 可以提供的功能:
* 返回排序過的graph data.
* 依據每一台機器的某一個監控項的time-series data計算 "mean","max","min" 三種計算結果數值
* 可以指定依照 mean 還是 max 或是 min 來排序
* 可以做正反序排序指定 desc / asc

[源碼解釋]:

 * 1 ~ 3行: 撰寫給提供給user外部指定參數的預設值.
 * (isNaN(v.Value)? 0 : v.Value) 為解決graph某些metric的 graph data value 為 "null" 值導致計算錯誤, 所以如果該值為NaN就將其轉換為 "0"

filepaht: `./js/top.js`

```
limit = (typeof limit == "undefined"? 3 : limit)
orderby = (typeof orderby == "undefined"? "desc" : orderby)
sortby = (typeof sortby == "undefined"? "Mean" : sortby)
t2 = _.map(input, function(res){
  res.Mean = 0
  res.Max = 0
  res.Min = 0
  if( res.Values.length == 0){
    return res
  }else{
    values = []
    mean = _.reduce(res.Values, function(sum,v){
      value = (isNaN(v.Value)? 0 : v.Value)
      values.push(value)
      return (sum+value)
    },0) / (res.Values.length === 0 ? 1 : res.Values.length)
    res.Mean = Math.round(mean,0)
    res.Max = Math.round(_.max(values), 0)
    res.Min = Math.round(_.min(values), 0)
    return res
  }
})

t3 = _.chain(t2).sortBy(function(res){

  if(orderby == "desc"){
    return - res[sortby]
  }else{
    return res[sortby]
  }

}).first(limit).value()

output = JSON.stringify(t3)
```

在設定檔加上對應的設置 `./conf/lambdaSetup.json`

```
{
    "funcation_name": "top",
    "file_path": "top.js",
    "params": ["limit:int", "orderby:string", "sortby:string"],
    "description": "找出topN的機器. orderby 調整正負排序. sortby指定比較基準數值 Mean/Max/Min"
}
```

[注意] 完成以上步驟之後就大功完成了, 但是當你撰寫完一個新的function且擺到query之後,需要重啟gin_http server 設定才可以生效.
