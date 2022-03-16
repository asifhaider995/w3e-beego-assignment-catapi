{{template "layout.tpl" . }}

{{define "styles" }}
    <style type="text/css">
        .main {
            display: flex;
            justify-content: center;
            align-items: center;
        }
        .header {
            background: center center;

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
            <h1 class="logo-2">Find a cat by Category and Image</h1>

            <div class="centered">
                <div>
                    <form id="query_form" class="form-horizontal form-well" role="form" action="/filtered-by-category" method="post">
                        <select name="category">
                            {{$cname := .categoryName}}
                            {{range $key, $val := .Categories}}
                                <option name="category" {{if eq $cname $val}} selected {{end}} value="{{$key}}">{{$val}}</option>
                            {{end}}
                        </select>
                        <select name="mime_type">
                            <option value="">Any</option>
                            <option value="jpg">Static (JPG) </option>
                            <option value="png">Static (PNG) </option>
                            <option value="gif">Animated</option>
                        </select>
                        <button type="submit">Search</button>
                    </form>

                </div>
                <div>
                    {{/*                <div>{{.infoLength}}</div>*/}}
                    {{if gt .infoLengthCat 2}}
                        <div class="results">
                            <img class="imgDiv" src="" alt="cat">
                            <h4 class="name"></h4>
                            <p class="description"></p>
                        </div>
                    {{else if eq .infoLengthCat 0}}
                        <div>
                            <h4>Please click the search button</h4>
                        </div>
                    {{else}}
                        <div>
                            <h4>No data are available for the Category {{.categoryName}}</h4>
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
            let info = JSON.parse({{.infoCat}})
            console.log(info[0])
            let url = info[0]["url"]
            u('.imgDiv').attr('src', url).attr('width', 350).attr('height', 250);
            // let xml = new XMLHttpRequest();
            //
            // xml.onload = function () {
            //     // console.log(JSON.parse(xml.response)[0].url)
            //     u('.imgDiv').attr('src', JSON.parse(xml.response).url).attr('width', 350).attr('height', 250);
            // }
            // xml.open("GET", "https://api.thecatapi.com/v1/images/"+img_ref)
            // xml.send()
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