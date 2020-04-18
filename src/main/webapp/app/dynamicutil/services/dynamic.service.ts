import {Injectable} from '@angular/core';
import {ActionInfo} from "app/dynamicutil/models/ActionInfo";
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {Observable} from "rxjs";
import {SERVER_API_URL} from "app/app.constants";
import * as CS from "app/dynamicutil/models/Constants";

@Injectable({
  providedIn: 'root'
})
export class DynamicService {

  constructor(private http: HttpClient) {

  }

  execute<T>(command: string, data: string): Observable<T> {
    const actionInfo: ActionInfo =
      {command, data};

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      })
    };

    // if (actionInfo.command === CS.GETSUBSYSTEMS) {
    //   return this.http.get<T>('assets/subsystems.json', httpOptions);
    // }

    // if (actionInfo.command === CS.GETSUBSYSTEMCOMMANDS) {
    //   return this.http.get<T>('assets/commands.json', httpOptions);
    // }

    // if (actionInfo.command === "lookupfakedata") {
    //   return this.http.get<T>('assets/lookupfakedata.json', httpOptions);
    // }


    return this.http.get<T>(SERVER_API_URL + 'api/' + command + data, httpOptions)
  }
}
