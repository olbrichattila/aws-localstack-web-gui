<form action="{{ route('file_upload') }}" method="post" enctype="multipart/form-data">
    <!-- @csrf -->
    <input type="file" name="file">
    <button type="submit">Upload</button>
</form>
