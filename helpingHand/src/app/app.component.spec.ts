import { TestBed, ComponentFixture } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { AppComponent } from './app.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';

import { DebugElement, NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatSlideToggleModule } from '@angular/material/slide-toggle'
import { MatButtonModule } from '@angular/material/button'
import { MatInputModule } from '@angular/material/input'; 
import { MatCardModule } from '@angular/material/card';
import {MatSelectModule} from '@angular/material/select'; 
import {MatTabsModule} from '@angular/material/tabs';
import { FormsModule } from '@angular/forms';
import { lastValueFrom } from 'rxjs'
import { By } from '@angular/platform-browser'

describe('AppComponent', () => {
  let debugElement: DebugElement;
  let fixture: ComponentFixture<AppComponent>;
  let app: AppComponent;
  let rightSideBar: HTMLElement;
  let images: HTMLElement;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        RouterTestingModule,
        HttpClientTestingModule,
        BrowserModule,
        AppRoutingModule,
        BrowserAnimationsModule,
        MatSlideToggleModule,
        MatButtonModule,
        HttpClientModule,
        FormsModule,
        MatCardModule,
        MatInputModule,
        MatTabsModule,
        MatSelectModule
      ],
      declarations: [
        AppComponent
      ],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AppComponent);
    app = fixture.componentInstance;
    let suits: string[] = ["spade", "club", "heart", "diamond"];
    for(let i = 0; i < 7; i++) {
      app.currentHand.push({
        Suit: suits[Math.floor(Math.random() * 3)],
        Val: Math.floor(Math.random() * 13) + 1,
        Index: i
      });
      app.currImgs.push("../assets/" + (app.currentHand[i].Val + 1) + app.currentHand[i].Suit + ".png");
    }
    fixture.detectChanges();
  });


  //First test to ensure that the app actually gets created

  it('should create the app', () => {
    expect(app).toBeTruthy();
  });

  it('should have cards equal to images', () => {
    const lengthCurrHand = app.currentHand.length;
    const lengthImgHand = app.currImgs.length;
    expect(lengthCurrHand).toEqual(lengthImgHand);
  });

  it('the hole cards should only contain two images', () => {
    const element: DebugElement[] = fixture.debugElement.queryAll(By.css('#cardImgs'));
    expect(element.length).toBe(2);
  })

  it('the community cards should only contain five images', () => {
    const element: DebugElement[] = fixture.debugElement.queryAll(By.css('#bottomcardImgs'));
    expect(element.length).toBe(5);
  })

  it('should not display selection screen when not flagged', () => {
    expect(fixture.debugElement.query(By.css('#displaySuitVal'))).toBeNull();
  })

  it('should display selection screen when flagged', () => {
    app.displaySuitVal = true;
    expect(fixture.debugElement.query(By.css('#displaySuitVal'))).toBeDefined();
  })

  it('should call addCard for each card in currentHand in removeAll', () => {
    spyOn(app, 'addCard')
    app.removeAll();
    expect(app.addCard).toHaveBeenCalledTimes(app.currentHand.length); 
  })

  it('should call addCard for each card in currentHand in randomizeAll', () => {
    spyOn(app, 'addCard')
    app.randomizeAll();
    expect(app.addCard).toHaveBeenCalledTimes(app.currentHand.length); 
  })

  it('setSuit should call addCard if random is passed in as a parameter', () => {
    spyOn(app, 'addCard')
    app.setSuit('random');
    expect(app.addCard).toHaveBeenCalled(); 
  })

  it('addCard should not be called if -1 is passed to setVal', () => {
    spyOn(app, 'addCard')
    app.setVal(-1);
    expect(app.addCard).not.toHaveBeenCalled(); 
  })

  it('displaySuit should swap the value of displaySuitVal', () => {
    app.displaySuitVal = false;
    app.displaySuit(1);
    expect(app.displaySuitVal).toBeTruthy(); 
  })

  it('right click handler function should call addCard', () => {
    spyOn(app, 'addCard')
    app.handleRightClick(-1, -1 ,'test', event);
    expect(app.addCard).toHaveBeenCalled(); 
  })



  
  
  /*

  it('should display new text when a card is added to currentHand', () => {
    const element: DebugElement[] = fixture.debugElement.queryAll(By.css('.rightsideloop'));
    expect(element.length).toBe(1);
  })

  it('should display an image when one is added to currImgs', () => {
     const element: DebugElement[] = fixture.debugElement.queryAll(By.css('.cardImgs'));
     expect(element.length).toBe(1);
  })

  it('should display the proper values of currentHand on the page', () => {
    const element: DebugElement[] = fixture.debugElement.queryAll(By.css('.rightsideloop'));
    element.forEach((obj:DebugElement, index:number) => {
      expect(obj.children[0].nativeElement.innerHTML.trim()).toEqual(app.currentHand[index].Suit + " " + app.currentHand[index].Val.toString());
    })
  })

  it('should display the proper image according to currImage', () => {
    const element =  fixture.debugElement.queryAll(By.css('.cardImgs'));
    element.forEach((obj:DebugElement, index:number) => {
      //Src is shown as http://localhost, meanwhile currImgs shows ..//assets, slice is used to remove the first few dots.
      expect(obj.children[0].children[0].nativeElement.src).toContain((app.currImgs[index]).slice(2));
    })
  })

  it('should display the proper image according to currentHand', () => {
    const element =  fixture.debugElement.queryAll(By.css('.cardImgs'));
    element.forEach((obj:DebugElement, index:number) => {
      expect(obj.children[0].children[0].nativeElement.src).toContain("/assets/" + (app.currentHand[index].Val) + app.currentHand[index].Suit + ".png");
    })
  })*/

  

  


  
});
