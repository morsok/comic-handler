import { TestBed } from '@angular/core/testing';
import { App } from './app';
import { ComicService } from './services/comic.service';
import { of } from 'rxjs';

describe('App', () => {
  const mockComicService = {
    getComics: () => of([])
  };

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [App],
      providers: [
        { provide: ComicService, useValue: mockComicService }
      ]
    }).compileComponents();
  });

  it('should create the app', () => {
    const fixture = TestBed.createComponent(App);
    const app = fixture.componentInstance;
    expect(app).toBeTruthy();
  });

  it(`should have the 'Comic Handler' title`, () => {
    const fixture = TestBed.createComponent(App);
    const app = fixture.componentInstance;
    expect(app['title']()).toEqual('Comic Handler');
  });

  it('should render title', () => {
    const fixture = TestBed.createComponent(App);
    fixture.detectChanges();
    const compiled = fixture.nativeElement as HTMLElement;
    // Assuming the title is rendered in a specific element, checking for existence of the text
    expect(compiled.textContent).toContain('Comic Handler');
  });
});
