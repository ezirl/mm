<form action="/kosmo/election/edit/" method="post" class="login">
    Title: <input name="title" type="text" value="{{ .election.Title }}">
    Link: <input name="link" type="text" value="{{ .election.Link}}">
    Site Name: <input type="text" name="sitename" value="{{.election.SiteName}}">
    Description: <input type="text" name="description" value="{{.election.Description}}">
    <input type="hidden" name="id" value="{{.election.ID}}">
    <button>update article</button>
</form>