using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace DatabaseAPI.Database.Models;

public class User(string name, string surname, string login, string password, string roleKey)
{
	[Key]
	[DatabaseGenerated(DatabaseGeneratedOption.Identity)]
	public int Id { get; set; }

	[Required]
	public string Name { get; set; } = name;

	[Required]
	public string Surname { get; set; } = surname;

	public int Balance { get; set; } = 0;

	[Required]
	public string Login { get; set; } = login;

	[Required]
	public string Password { get; set; } = password;

	[Required]
	public string RoleKey { get; set; } = roleKey;

	[ForeignKey(nameof(RoleKey))]
	public Role Role { get; set; }
}
