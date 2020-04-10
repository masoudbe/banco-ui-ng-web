import { Injectable } from '@angular/core';
import {BehaviorSubject, Observable} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class PresenterStoreService {

  private presenters: string[] = [];
  private subject = new BehaviorSubject<string[]>([]);
  presenters$: Observable<string[]> = this.subject.asObservable();

  constructor() { }

  addPresenter(presenterName: string){
    this.presenters.push(presenterName);
  }

  removePresenter(presenterName: string){
    const index = this.presenters.indexOf(presenterName, 0);
    if (index > -1) {
      this.presenters.splice(index, 1);
    }
  }
}
