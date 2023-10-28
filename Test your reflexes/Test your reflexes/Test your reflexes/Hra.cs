using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace Test_your_reflexes
{
    public partial class Hra : Form
    {
        private int pocetPokusu;
        private int velikostBoxu;
        private Graphics G;
        private int x;
        private int y;
        private int pocitadlo;
        private int pocetTrefenych;
        private int max_obrazx;
        private int max_obrazy;
        private int stredx;
        private int stredy;
        private string presnost;
        public Hra(int p, int t, int v, int x, int y)
        {
            InitializeComponent();
            this.pocetPokusu = p;
            this.velikostBoxu = v;
            this.max_obrazx = x;
            this.max_obrazy = y;
            GoFullscreen(true);
            this.G = pictureBox1.CreateGraphics();
            this.pocitadlo = 0;
            this.pocetTrefenych = 0;
            this.timer1.Interval = t;
            this.presnost = "";
        }
        public Hra()
        {
            this.velikostBoxu = 0;
            InitializeComponent();
            GoFullscreen(true);
            this.G = pictureBox1.CreateGraphics();
            
        }

        private void pictureBox1_MouseClick(object sender, MouseEventArgs e)
        {
            bool test = false;
            if (velikostBoxu == 0)
            {
                this.G.FillRectangle(Brushes.Green, 1, 1, 1918, 1078);
                this.G.FillRectangle(Brushes.Red, 1, 1, 1678, 1048);
                this.G.FillRectangle(Brushes.Blue, 1, 1, 998, 498);
                test = true;
            }


            int ix = e.X;
            int iy = e.Y;

            if (ix < x + velikostBoxu && ix > x)
            {
                //vzdálenost od středu by byla přesnost
                if (iy < y + velikostBoxu && iy > y)
                {
                    pocetTrefenych++;
                    this.Text = pocetTrefenych + "/" + pocetPokusu;
                    this.G.Clear(Color.Green);
                    this.x = -1000;
                    this.y = -1000;
                    
                }
            }
            else if (test == false && this.timer1.Enabled == true)
            {
                this.G.Clear(Color.Red);
                this.x = -1000;
                this.y = -1000;
            }
            else if(test == false && this.timer1.Enabled == false)
            {
                this.G.Clear(Color.Yellow);
            }

            if (pocitadlo == 0)
            {
                this.timer1.Enabled = true;

            }
        }

        private void pictureBox1_MouseDoubleClick(object sender, MouseEventArgs e)
        {
            if (pocetTrefenych == 0)
            {
                if (velikostBoxu != 0) { MessageBox.Show("Tvoje skóre se po zaokrouhlení rovná nule."); }
                        
                this.Hide();
                
            }
            else if (pocetTrefenych == pocitadlo)
            {
                this.Hide();
                MessageBox.Show("Plný počet.");
            }
            else
            {
                this.Hide();
                MessageBox.Show(pocetTrefenych + "/" + pocetPokusu);
            }
        }
        private void GoFullscreen(bool fullscreen)
        {
            if (fullscreen)
            {
                this.WindowState = FormWindowState.Normal;
                this.FormBorderStyle = System.Windows.Forms.FormBorderStyle.None;
                this.Bounds = Screen.PrimaryScreen.Bounds;
            }
            else
            {
                this.WindowState = FormWindowState.Maximized;
                this.FormBorderStyle = System.Windows.Forms.FormBorderStyle.Sizable;
            }
        }

        private void timer1_Tick(object sender, EventArgs e)
        {
            if (velikostBoxu != 0)
            {
                if (pocitadlo < pocetPokusu)
                {
                    Random R = new Random();
                    this.x = R.Next(0, this.max_obrazx - this.velikostBoxu);
                    this.y = R.Next(0, this.max_obrazy - this.velikostBoxu);
                    if (this.velikostBoxu % 2 == 0)
                    {
                        this.stredx = (this.x + this.velikostBoxu) / 2+this.x;
                        this.stredy = (this.y + this.velikostBoxu) / 2+this.y;
                    }
                    else
                    {
                        this.stredx = (this.x + this.velikostBoxu+1) / 2+this.x;
                        this.stredy = (this.y + this.velikostBoxu+1) / 2+this.y;
                    }
                    this.G.Clear(Color.White);
                    this.G.FillRectangle(Brushes.Black, x, y, velikostBoxu, velikostBoxu);
                    pocitadlo++;
                }
                else
                {
                    this.x = 0;
                    this.y = 0;
                    timer1.Enabled = false;
                    this.G.Clear(Color.White);
                    

                }
            }
        }
    }
}
