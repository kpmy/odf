import 'dart:html';
import 'package:crypto/crypto.dart';
import 'dart:typed_data';
import 'dart:js';
import 'dart:convert';

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

  w.listen((_m){
    print(_m.data);
    var msg = JSON.decode(_m.data);
    switch(msg["Type"]){
      case "init":
        print("worker initialized, sending responce...");
        w.postMessage(JSON.encode({"Type": "init"}));
        break;
      case "data":
        print("data received");
        Uint8List data = new Uint8List.fromList(CryptoUtils.base64StringToBytes(msg["Data"]));
        context.callMethod("saveAs",  [new Blob([data], "application/octet-stream"), "report.odf"]);
        break;
      default: throw new ArgumentError(msg["Type"]);
    }
  });

  querySelector("#do-demo").onClick.listen((m){
    w.postMessage(JSON.encode({"Type": "get", "Param": "demo"}));
  });

  querySelector("#do-report").onClick.listen((m){
    w.postMessage(JSON.encode({"Type": "get", "Param": "report"}));
  });

}
