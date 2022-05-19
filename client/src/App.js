import './App.css';
import { ClientReadableStream, } from 'grpc-web';

import { UpvoteServiceClient, } from "./gen/proto/upvote/v1/upvote_grpc_web_pb.js";
import { WatchBookRequest } from './gen/proto/upvote/v1/upvote_pb.js'


const client = new UpvoteServiceClient('http://localhost:8080', null, null);

let stream = ClientReadableStream


var req = new WatchBookRequest();
req.setTitle("Clean Code");


stream = client.watchBook(req).on('data', (response) => {
  let bookTitle = response.getTitle();
  let bookAuthor = response.getAuthor();
  let bookUpvotes = response.getLikes();
  console.log(bookTitle, bookAuthor, bookUpvotes);
}).on('status', (status) => {
  console.log(status);
}).on('end', () => {
  console.log('end');
});


function App() {

  return (

    <div className="App">

    </div>
  );
}

export default App;
