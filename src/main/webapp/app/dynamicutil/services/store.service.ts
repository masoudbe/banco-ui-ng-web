import {Injectable} from '@angular/core';
import {BehaviorSubject} from "rxjs";
import {CommandInfo} from "app/dynamicutil/models/CommandInfo";

@Injectable({
  providedIn: 'root'
})
export class StoreService {

  private commandsSubject = new BehaviorSubject<CommandInfo[]>([]);
  commands$ = this.commandsSubject.asObservable();

  private sidebarToggleSubject = new BehaviorSubject<void>(undefined);
  toggleSideBar$ = this.sidebarToggleSubject.asObservable();

  constructor() {

  }

  addPresenter(ci: CommandInfo): void {
    const cis = this.commandsSubject.getValue();
    const newCis = cis.slice(0);
    newCis.push(ci);

    this.commandsSubject.next(newCis);
  }

  toggleSideBar(): void {
    this.sidebarToggleSubject.next();
  }

  removePresenter(ci: CommandInfo): void {
    const prs = this.commandsSubject.getValue();
    const newPrs = prs.slice(0);
    const index = newPrs.indexOf(ci, 0);
    if (index > -1) {
      newPrs.splice(index, 1);
    }

    this.commandsSubject.next(newPrs);
  }
}
