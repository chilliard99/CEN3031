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
//Errors in this file are from conflict between cypress & jasmine, tests run fine.
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
    app.currentProb.push({
      Handname: "High Card",
      Prob: 0
    });
    app.currentProb.push({
      Handname: "One Pair",
      Prob: 0
    });
    app.currentProb.push({
      Handname: "Two Pair",
      Prob: 0
    });
    app.currentProb.push({
      Handname: "Three of a Kind",
      Prob: 0
    });
    app.currentProb.push({
      Handname: "Straight",
      Prob: 0
    });
    app.currentProb.push({
      Handname: "Flush",
      Prob: 0
    });
    app.currentProb.push({
      Handname: "Full House",
      Prob: 0
    });
    app.currentProb.push({
      Handname: "Four of a Kind",
      Prob: 0
    });
    app.currentProb.push({
      Handname: "Straight Flush",
      Prob: 0
    });
    app.currentProb.push({
      Handname: "Royal Flush",
      Prob: 0
    });
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

  //Added in Sprint 4

  it('toggle dark mode should should switch darkMode', () => {
    app.darkMode = false;
    app.toggleDarkMode();
    expect(app.darkMode).toBeTruthy(); 
  })

  it('darkMode should not be displayed when darkmode is not activated', () => {
    app.darkMode = false;
    expect(fixture.debugElement.query(By.css('.holecardsdark'))).toBeNull();
    expect(fixture.debugElement.query(By.css('.simulationNumsDark'))).toBeNull();
    expect(fixture.debugElement.query(By.css('.moreInfoButtonDark'))).toBeNull();
    expect(fixture.debugElement.query(By.css('.communitycardsdark'))).toBeNull();
  })

  it('darkMode should be displayed when darkmode is activated', () => {
    app.darkMode = true;
    expect(fixture.debugElement.query(By.css('.holecardsdark'))).toBeDefined();
    expect(fixture.debugElement.query(By.css('.simulationNumsDark'))).toBeDefined();
    expect(fixture.debugElement.query(By.css('.moreInfoButtonDark'))).toBeDefined();
    expect(fixture.debugElement.query(By.css('.communitycardsdark'))).toBeDefined();
  })

  it('currentProb should have only 10 members (number of possible hands)', () => {
    expect(app.currentProb.length).toBe(10);
  })

  it('currentProb should only have 0s in prob', () => {
    let probSum = 0;
    for(let i = 0; i < app.currentProb.length; i++) {
      probSum += app.currentProb[i].Prob
    }
    expect(probSum).toBe(0);
  })

  it('checkForRepeats should return true if there is a repeat', () => {
    app.currentHand[0].Suit = "Spade";
    app.currentHand[0].Val = 0;
    expect(app.checkForRepeats(1,"Spade")).toBeTruthy();
  })

  it('changeTab should set the tab', () => {
    app.changeTab(0);
    expect(app.selectedIndex).toBe(0);
    app.changeTab(1);
    expect(app.selectedIndex).toBe(1);
  })

  it('Passing in an illegal value into setVal should not call addCard', () => {
    spyOn(app, 'addCard')
    app.currentHand[0].Suit = "Spade";
    app.currentHand[0].Val = 0;
    app.Suit = "Spade";
    app.setVal(1);
    expect(app.addCard).not.toHaveBeenCalled();  
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
