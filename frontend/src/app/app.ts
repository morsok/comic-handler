import { Component, signal, inject, OnInit, ChangeDetectionStrategy } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { ComicService, Comic } from './services/comic.service';
import { ThemeService } from './services/theme.service';
import { ComicCardComponent } from './components/comic-card.component';

@Component({
  selector: 'app-root',
  imports: [HttpClientModule, ComicCardComponent],
  templateUrl: './app.html',
  styleUrl: './app.sass',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class App implements OnInit {
  private readonly comicService = inject(ComicService);
  private readonly themeService = inject(ThemeService);
  
  protected readonly title = signal('Comic Handler');
  protected readonly comics = signal<Comic[]>([]);
  protected readonly loading = signal(true);
  protected readonly error = signal<string | null>(null);
  protected readonly isDarkMode = this.themeService.isDarkMode;

  ngOnInit(): void {
    this.loadComics();
  }

  protected toggleTheme(): void {
    this.themeService.toggleTheme();
  }

  protected loadComics(): void {
    this.loading.set(true);
    this.error.set(null);
    
    this.comicService.getComics().subscribe({
      next: (comics) => {
        this.comics.set(comics);
        this.loading.set(false);
      },
      error: (err) => {
        this.error.set('Failed to load comics. Please try again.');
        this.loading.set(false);
        console.error('Error loading comics:', err);
      }
    });
  }

  protected retryLoad(): void {
    this.loadComics();
  }
}
