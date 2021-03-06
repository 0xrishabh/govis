package util

func Template() []byte {
	template := `<html>
    <head>
        <title>Recon Dashboard</title>
        <link href="https://fonts.googleapis.com/css2?family=Kaushan+Script&family=Merienda+One&display=swap" rel="stylesheet">
        <script src="https://code.jquery.com/jquery-3.5.1.min.js" integrity="sha256-9/aliU8dGd2tb6OSsuzixeV4y/faTqgFtohetphbbj0=" crossorigin="anonymous"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery-modal/0.9.1/jquery.modal.min.js"></script>
        <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/jquery-modal/0.9.1/jquery.modal.min.css" />
        <script src="./data.js"></script>
        <style>
            html {
                scroll-behavior: smooth;
            }
            body{
                margin: 0;
                font-family: 'Roboto', sans-serif;

            }
            .container{
                display: flex;
                flex-wrap: wrap;
                justify-content: space-between;
                max-width: 100vw;
                padding: 5px;
                margin-top: 40px;
            }
            .wrap{
                width: 100%;
                display: grid;
                grid-template-columns: 1fr 2fr;
                padding: 5px;
                background-color: #fff;
                margin-top: 5px;
                box-shadow: 0px 0px 5px 0px #2c7eff;
            }
            .thumbnail{
                width: 400px;
                max-height: 250px;
                cursor: pointer;
                border-right: 3px solid #2c7eff;
            }
            .hiddenjs{
                display: none;
            }
            table {
                border-collapse: collapse;
                width: 100%;
            }
            .dot {
                height: 15px;
                width: 15px;
                border-radius: 50%;
                display: inline-block;
            }
            .dot.green{
                background-color: green;
            }
            .dot.red{
                background-color: red;
            }
            th, td {
                padding: 8px;
                text-align: left;
                text-align: center;
                font-family: 'Merienda One', cursive;
            }
            tr:hover {background-color:#f5f5f5;}
            button {
                border: none;
                text-align: center;
                text-decoration: none;
                display: inline-block;
                font-size: 16px;
                margin: 4px 2px;
                transition-duration: 0.4s;
                cursor: pointer;
                background-color: white;
                color: black;
                border: 2px solid #008CBA;
            }
            button:hover {
                background-color: #008CBA;
                color: white;
            }
            header{
                width: 100vw;
                overflow: hidden;
                padding: 5px;
                font-size: 24px;
                background-color: black;
                color: #fff;
                text-align: center;
                font-family: 'Kaushan Script', cursive;
                position: fixed;
                top: 0;
            }
        </style>

    </head>


    <body>
        <header>Govis</header>
        <div class="container">
        </div>
    </body>
    <script>
         function addData(ele){
            let url = ele.Url
            let title = ele.Title || "No Title"
            let statusCode = ele.StatusCode
            if (statusCode == 200){url += "  <span class='dot green'></span>"}
            else{url += "  <span class='dot red'></span>"}
            let imageSrc = ele.ImgPath
            let jsFiles = "";
            ele.JsUrlsList.forEach(function(x){jsFiles+=x+"\n"})
            let template = ` + "`" + `<div class="wrap" id="unit">
                <img class="thumbnail zoomable" src="${imageSrc}">
                <div class="info">
                    <div class="hiddenjs" class="files">${jsFiles}</div>
                    <table cellspacing="5">
                        <tr><th colspan=2>${url}</th></tr>
                        <tr><td class="subhead">Status Code</td><td>${statusCode}</td></tr>
                        <tr><td class="subhead">Title</td><td>${title}</td></tr>
                        <tr><td class="subhead">JsFiles</td><td><button class="jsLinks">${ele.JsUrlsList.length} js file found</button></td></tr>
                        </table>
                </div>
            </div>` + "`" + `
            $(".container").append(template)
        }
        Db.forEach(addData);

        $('.jsLinks').click(function(event){navigator.clipboard.writeText($(event.target).parent().parent().parent().parent().siblings()[0].innerText);alert('Javascript files copied to clipboard')})

        $('img.zoomable').css({cursor: 'pointer'}).click(function () {
            var img = $(this);
            var bigImg = $('<img />').css({
                'max-width': '100%',
                'max-height': '100%',
                'display': 'inline'
            });
            bigImg.attr({
                src: img.attr('src'),
                alt: img.attr('alt'),
                title: img.attr('title')
            });
            var over = $('<div />').text(' ').css({
                'height': '100%',
                'width': '100%',
                'background': 'rgba(0,0,0,.82)',
                'position': 'fixed',
                'top': 0,
                'left': 0,
                'opacity': 0.0,
                'cursor': 'pointer',
                'z-index': 9999,
                'text-align': 'center'
            }).append(bigImg).bind('click', function () {
                $(this).fadeOut(100, function () {
                $(this).remove();
                });
            }).insertAfter(this).animate({
                'opacity': 1
            }, 300);
        });
    </script>
</html>
`
	return []byte(template)
}
