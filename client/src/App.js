import './App.css';
import * as grpcWeb from 'grpc-web';

import { UpvoteServiceClient } from "./gen/proto/upvote/v1/upvote_grpc_web_pb.js";
import { ListBooksRequest } from './gen/proto/upvote/v1/upvote_pb.js'


const client = new UpvoteServiceClient('http://localhost:8080', null, null);

var request = new ListBooksRequest();

request.setTitle("Clean Code");

var stream = client.listBooks(request, {});

stream.on('data', (response) => {
  console.log(response.getTitle());
});

stream.on('error', (error) => {
  console.log(error);
});

function App() {

  return (

    <div className="App">

    </div>
  );
}

export default App;
