using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace DatabaseAPI.Database.Models;

public class Merch(string name, string description, int price)
{
	[Key]
	[DatabaseGenerated(DatabaseGeneratedOption.Identity)]
	public int Id { get; set; }
	[Required]
	public string Name { get; set; } = name;
	[Required]
	public string Description { get; set; } = description;
	[Required]
	public int Price { get; set; } = price;
}
