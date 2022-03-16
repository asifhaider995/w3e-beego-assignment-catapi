{{template "layout.tpl" . }}

{{define "styles" }}
    <style type="text/css">
        .main {
            display: flex;
            justify-content: center;
            align-items: center;
        }
        .centered {
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: space-evenly;
            min-height: 50px;
        }
        .results {
            display: flex;
            justify-content: space-between;
        }
        .imgDiv {
            margin: 1.5rem;
            justify-content: center;
            display: flex;
            align-items: center;
            border: 3px solid black;
            box-shadow: #444444;
            border-radius: 5px;
        }
        .name {
            margin-bottom: 5px;
        }
    </style>

{{end}}

{{define "content"}}
<div class="main">
    <main>
        <h1 class="logo-2">Find a cat by Breed</h1>
        <div class="centered">
            <div>
                <form id="query_form" class="form-horizontal form-well" role="form" action="/filtered-by-breed" method="post">
                    <select name="breed">
                        {{$bname := .breedName}}
                        {{range $key, $val := .Breeds}}
                            <option name="breed" {{if eq $bname $val}} selected {{end}} value="{{$key}}">{{$val}}</option>
                        {{end}}
                    </select>
                    <button type="submit">Search</button>
                </form>

            </div>
            <div>
{{/*                <div>{{.infoLength}}</div>*/}}
                    {{if gt .infoLength 2}}
                        <div class="results">
                            <img class="imgDiv" src="" alt="cat">
                            <h4 class="name"></h4>
                            <p class="description"></p>
                        </div>
                    {{else if eq .infoLength 0}}
                        <div>
                            <h4>Please click the search button</h4>
                        </div>
                    {{else}}
                        <div>
                            <h4>No data are available for the breed {{.breedName}}</h4>
                        </div>
                    {{end}}
            </div>
        </div>
    </main>
</div>
{{end}}

{{define "scripts"}}
<script>
    console.log("Info",{{.info}})
    try {
        let info = JSON.parse({{.info}})
        console.log(info[0])
        let img_ref = info[0]["reference_image_id"]
        // u('.name').add(info[0].name)
        // u('.description').add(info[0].description)
        let xml = new XMLHttpRequest();

        xml.onload = function () {
            // console.log(JSON.parse(xml.response)[0].url)
            u('.imgDiv').attr('src', JSON.parse(xml.response).url).attr('width', 350).attr('height', 250);

        }
        xml.open("GET", "https://api.thecatapi.com/v1/images/"+img_ref)
        xml.send()
    } catch (e) {
        console.log("Error", e)
    }





</script>
{{end}}

{{define "website"}}
    {{.Website}}
{{end}}

{{define "email"}}
    {{.Email}}
{{end}}