# TODO

- [ ] Create a getFileInfo function
- [ ] The function will accept an http client and a link
- [ ] The function will get the file total download size, guess the file name, check if the file accepts the download in chunks using the accept-ranges header, and finally get the mime type then use the mime stand package to get an extension out of it if possible
- [ ] If the above function doesn't get the filename, ask the user to enter the filename and also remember to show the mime type when asking for filename
- [ ] Create function in charge of calculating the optimum chunckSize using the total downloadSize of the file
- [ ] Create function for synchronous file download that will use io.Copy to copy the response directly to a file until it's complete
- [ ] Create a function for downloading files in chunks
- [ ] Show download progress
- [ ] Use go routines to download multiple chunks concurrently
- [ ] Show concurrent download progress
- [ ] Add pause and resume capabilities
- [ ] Add support youtube videos
- [ ] Add support twitter videos
- [ ] Add support tiktok videos
- [ ] Add support instagram videos
- [ ] Add support facebook videos
- [ ] Add support for torrent files
- [ ] Add a rest API that would be consumed by a frontend client