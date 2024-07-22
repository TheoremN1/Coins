using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace DatabaseAPI.Database.Models;

public class Achievement(string name, string description, int reward)
{
	[Key]
	[DatabaseGenerated(DatabaseGeneratedOption.Identity)]
	public int Id { get; set; }
	[Required]
	public string Name { get; set; } = name;
	[Required]
	public string Description { get; set; } = description;
	[Required]
	public int Reward { get; set; } = reward;
}
