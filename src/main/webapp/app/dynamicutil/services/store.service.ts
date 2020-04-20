import {Injectable} from '@angular/core';
import {BehaviorSubject} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class StoreService {

  private presentersSubject = new BehaviorSubject<string[]>([]);
  presenters$ = this.presentersSubject.asObservable();

  private sidebarToggleSubject = new BehaviorSubject<void>(undefined);
  toggleSideBar$ = this.sidebarToggleSubject.asObservable();

  constructor() {

  }

  toggleSideBar(): void {
    this.sidebarToggleSubject.next();
  }

  addPresenter(commandCode: string): void {
    const prs = this.presentersSubject.getValue();
    const newPrs = prs.slice(0);
    newPrs.push(commandCode);

    this.presentersSubject.next(newPrs);
  }

  removePresenter(presenterName: string): void {
    const prs = this.presentersSubject.getValue();
    const newPrs = prs.slice(0);
    const index = newPrs.indexOf(presenterName, 0);
    if (index > -1) {
      newPrs.splice(index, 1);
    }

    this.presentersSubject.next(newPrs);
  }
}
