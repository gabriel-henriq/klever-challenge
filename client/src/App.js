import './App.css';

import { UpvoteServiceClient, } from "./gen/proto/upvote/v1/upvote_grpc_web_pb.js";
import { WatchBookRequest } from './gen/proto/upvote/v1/upvote_pb.js'


const client = new UpvoteServiceClient('http://localhost:8080', null, null);


var req = new WatchBookRequest();
req.setTitle("Clean Code");
let updatedBook = {}

client.watchBook(req).on('data', (response) => {
  updatedBook = {
    title: response.getTitle(),
    author: response.getAuthor(),
    likes: response.getLikes()
  }
  console.log(updatedBook)
})


function App() {

  return (
    <div className="App">
      <div>
        <h1>Book</h1>
      </div>
    </div>
  );
}

export default App;
