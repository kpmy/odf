import 'dart:html';

class OdfWorker{

  Worker inner;

  OdfWorker(){
    this.inner = new Worker("demo3.js");
    this.inner.onMessage.listen((m){
      print("worker initialized, sending responce...");
      this.inner.postMessage({'Typ': 'init'});
    });
  }
}

void main() {
  new OdfWorker();
}
