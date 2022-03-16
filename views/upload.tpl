{{template "layout.tpl".}}

{{define "styles" }}
    <style type="text/css">
        .upload {
            display: flex;
            justify-content: center;
            align-items: center;
            margin: 15px 0;
        }
        .form {
            display: flex;
            justify-content: space-evenly;
            align-items: center;
        }
    </style>

{{end}}

{{define "content"}}
    <h1 class="logo">Upload a Cat photo</h1>
    <div class="description">
        Upload the photo of your favorite cat
    </div>
    <div class="upload">
        <form name="form" id="upload_form" enctype="multipart/form-data">
            <input type="file" name="file"/>
            <button type="submit">Upload File</button>
        </form>
    </div>

{{end}}

{{define "website"}}
    {{.Website}}
{{end}}

{{define "email"}}
    {{.Email}}
{{end}}

{{define "scripts"}}
    <script>
        console.log("Hello World!")
        // u('.form').handle('change', function (e) {
        //     e.preventDefault()
        //     console.log(e)
        // })
        u('form').handle('submit', function (e) {
            e.preventDefault()
            let file_input = document.getElementsByName('file')

            console.log(file_input[0].files[0])
            let fd = new FormData()
            fd.append("file", file_input[0].files[0])
            // const pl = {
            //     "file": file_input[0].files[0]
            // }
            const url = "https://api.thecatapi.com/v1/images/upload"
            const api_key = "dcd005e2-35a8-45ec-843d-fcdfac3a4869"
            let xml = new XMLHttpRequest();
            xml.onload = () => {
                console.log(xml.responseText)

            }
            xml.open("POST", url, true);
            xml.setRequestHeader('Access-Control-Allow-Origin', "*");
            // xml.setRequestHeader('Content-Type', "multipart/form-data");
            // xml.setRequestHeader('Content-Type', "multipart/form-data;boundary=----WebKitFormBoundaryyrV7KO0BoCBuDbTL");
            // xml.setRequestHeader('Content-Type', "text/plain");
            xml.setRequestHeader('Mime-Type', "image/jpeg");
            xml.setRequestHeader('x-api-key', api_key);
            xml.send(fd)
        })
    </script>
{{end}}