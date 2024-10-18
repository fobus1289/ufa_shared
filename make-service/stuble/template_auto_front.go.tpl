<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>{{getTitle}}</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <style>
        body {
            font-family: Arial, sans-serif;
            padding: 20px;
        }
        h1 {
            color: #2c3e50;
        }
        .api-path {
            background-color: #ecf0f1;
            padding: 10px;
            margin-bottom: 15px;
        }
        .endpoint-details {
            display: block; 
            margin-top: 10px;
            padding: 10px;
            background-color: #f5f5f5;
            border: 1px solid #ddd;
        }
    </style>
</head>

<body>

    {{range $path, $pathItem := getPaths}}
    <div class="api-path">
        <h2>{{$path}}</h2>
        {{range $method, $operation := $pathItem}}
        <div id="{{$path}}" class="endpoint-details">
            <h5>{{$method}}</h5>
        </div>
        {{end}}
    </div>
    {{end}}
</body>
</html>
