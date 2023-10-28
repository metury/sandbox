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
    public partial class Form1 : Form
    {
        
        
        public Form1()
        {
            InitializeComponent();
            this.textBox1.Text = "20";
            this.textBox2.Text = "1";
            this.textBox3.Text = "50";
            this.comboBox1.Text = "zelená";
            this.comboBox1.SelectedItem = "zelená";
            
        }

        private void button1_Click(object sender, EventArgs e)
        {
            try
            {
                double x = double.Parse(textBox2.Text) * 1000;
                int y = Convert.ToInt32(x);
                if (y <= 0 || int.Parse(textBox1.Text) <= 0 || int.Parse(textBox3.Text) <= 0)
                {
                    MessageBox.Show("Zadej kladné číslo!", "CHYBA");

                }
                else if (int.Parse(textBox3.Text) > getY())
                {
                    MessageBox.Show("Takový čtverece se tam nevejde!", "CHYBA");
                }
                else
                {
                    Hra H = new Hra(int.Parse(textBox1.Text), y, int.Parse(textBox3.Text), getX(), getY()); ;
                    H.Show();
                }
            }
            catch 
            {
                MessageBox.Show("Zřejmě si zadal něco špatně.", "ERROR"); 
            }
            
            
            
        }

        

        private void button2_Click_1(object sender, EventArgs e)
        {
            Hra H = new Hra();
            H.Show();
        }

        private void comboBox1_SelectedIndexChanged(object sender, EventArgs e)
        {

        }
        private int getX()
        {
            if (comboBox1.SelectedItem == "červená")
            {
                return (1680);
            }
            if (comboBox1.SelectedItem == "zelená")
            {
                return (1920);
            }
            else { return (1000); }
        }
        private int getY()
        {
            if (comboBox1.SelectedItem == "červená")
            {
                return (1050);
            }
            if (comboBox1.SelectedItem == "zelená")
            {
                return (1080);
            }
            else { return (500); }
        }

        private void button3_Click(object sender, EventArgs e)
        {
            MessageBox.Show("Po zapnutí hry jedním kliknutím hru zapnete a dvojklikem ukončíte. Abyste zjistili velikost obrazu, tak zkuste test a zvolte barvu obdélníku, který vidíte celý. Jinak test funguje stejně jako hra, takže klik a dvojklik. Pokud chcete nastavit timer na desetinný čísla, tak zadejte s desetinnou tečkou. Velikost boxu je jedna strana boxu v pixelech", "INFO");
        }
    }
}
