# open lambda query

![](icon.png)

這是一個特製過的Query版本, 可以在Golang撰寫 Javascript Lambda function 處理資料畫圖

### Setting

```
cp cfg.example.json cfg.json
更新 cfg.json
```

* cfg.json 說明
  * root_dir: 設定working directory絕對路徑
  * graph: 設置Open-falcon Graph 所在位置,供查詢繪圖資料時使用
  * db: falcon_portal db connection 參數
  * graph: graph db connection 參數

### 完成以上設定

```
cd $working_directory
go build
./open_lambda_query
```
### 目前支援功能
```
[
  {
    "funcation_name": "avgCompare",
    "file_path": "avgCompare.js",
    "params": ["cond:string"],
    "description": "找出高/低於平均值的機器"
  },
  {
    "funcation_name": "top",
    "file_path": "top.js",
    "params": ["limit:int", "orderby:string", "sortby:string"],
    "description": "找出topN的機器. orderby 調整正負排序. sortby指定比較基準數值 Mean/Max/Min"
  },
  {
    "funcation_name": "sumAll",
    "file_path": "sumAll.js",
    "params": ["aliasName:string"],
    "description": "加總所有的數值"
  },
  {
    "funcation_name": "limit",
    "file_path": "limit.js",
    "params": ["limit:int"],
    "description": "只拿出前幾筆查條件來畫圖"
  },
  {
    "funcation_name": "topDiff",
    "file_path": "topDiff.js",
    "params": ["limit:int", "orderby:string", "sortby:string"],
    "description": "找出兩點間成長幅度最高/低的metric, orderby 調整正負排序. sortby指定比較基準數值 Mean/Max/Min"
  },
  {
    "funcation_name": "aliasName",
    "file_path": "aliasName.js",
    "params": ["name:string"],
    "description": "自定義顯示線圖名稱"
  }
]
```
![](./grafana4_lambda.png)
