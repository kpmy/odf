import 'dart:html';
import 'package:crypto/crypto.dart';
import 'dart:typed_data';
import 'dart:js';

class OdfWorker{

  Worker inner;

  void postMessage(message){
    this.inner.postMessage(message);
  }

  void listen(handler){
    this.inner.onMessage.listen(handler);
  }

  OdfWorker(){
    this.inner = new Worker("demo3.js");
  }
}

void main() {
  var w = new OdfWorker();

  w.listen((m){
    switch(m.data["Typ"]){
      case "init":
        print("worker initialized, sending responce...");
        w.postMessage({'Typ': 'init'});
        break;
      case "data":
        print("data received");
        Uint8List data = new Uint8List.fromList(CryptoUtils.base64StringToBytes(m.data["Data"]));
        context.callMethod("saveAs",  [new Blob([data], "application/octet-stream"), "report.odf"]);
        break;
      default: throw new ArgumentError(m.data["Typ"]);
    }
  });

  querySelector("#do-demo").onClick.listen((m){
    w.postMessage({'Typ': 'get', 'Param': 'demo'});
  });

  querySelector("#do-report").onClick.listen((m){
    w.postMessage({'Typ': 'get', 'Param': 'report'});
  });

}
