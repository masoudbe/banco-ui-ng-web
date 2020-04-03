import {Injectable} from '@angular/core';
import {ActionInfo} from "app/shared/dynamicutil/models/ActionInfo";
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {Observable} from "rxjs";
import {SERVER_API_URL} from "app/app.constants";

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

    return this.http.post<T>(SERVER_API_URL + 'execute', actionInfo, httpOptions)
  }
}
